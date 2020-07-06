package job

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/filanov/bm-inventory/internal/common"
	logutil "github.com/filanov/bm-inventory/pkg/log"
	"github.com/filanov/bm-inventory/restapi/operations/installer"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	batch "k8s.io/api/batch/v1"
	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const kubeconfigPrefix = "generate-kubeconfig"

// [TODO] need to find more generic way to set the openshift release image
//https://mirror.openshift.com/pub/openshift-v4/clients/ocp-dev-preview/4.5.0-0.nightly-2020-05-21-015458/
const OverrideImageName = "quay.io/openshift-release-dev/ocp-release-nightly@sha256:a9f7564e0f2edef2c15cc1da699ebd1d11f5acd717c3668940848b3fed0d13c7"

// [TODO]  make sure that we use openshift-installer from the release image, otherwise the KubeconfigGenerator image must be updated here per openshift version

type ISOGenerator interface {
	GenerateISO(ctx context.Context, cluster common.Cluster, jobName string, imageName string, ignitionConfig string) *installer.GenerateClusterISOInternalServerError
}

type InstallConfigGenerator interface {
	GenerateInstallConfig(ctx context.Context, cluster common.Cluster, cfg []byte) error
}

type ISOInstallConfigGenerator interface {
	ISOGenerator
	InstallConfigGenerator
}

//go:generate mockgen -source=job.go -package=job -destination=mock_job.go
type API interface {
	// Create k8s job
	Create(ctx context.Context, obj runtime.Object, opts ...client.CreateOption) error
	// Monitor k8s job return error in case job fails
	Monitor(ctx context.Context, name, namespace string) error
	// Delete k8s job
	Delete(ctx context.Context, name, namespace string) error
	ISOGenerator
	InstallConfigGenerator
}

type Config struct {
	MonitorLoopInterval time.Duration `envconfig:"JOB_MONITOR_INTERVAL" default:"500ms"`
	RetryInterval       time.Duration `envconfig:"JOB_RETRY_INTERVAL" default:"1s"`
	RetryAttempts       int           `envconfig:"JOB_RETRY_ATTEMPTS" default:"30"`
	ImageBuilder        string        `envconfig:"IMAGE_BUILDER" default:"quay.io/ocpmetal/installer-image-build:latest"`
	ImageBuilderCmd     string        `envconfig:"IMAGE_BUILDER_CMD" default:"echo hello"`
	Namespace           string        `envconfig:"NAMESPACE" default:"assisted-installer"`
	S3EndpointURL       string        `envconfig:"S3_ENDPOINT_URL" default:"http://10.35.59.36:30925"`
	S3Bucket            string        `envconfig:"S3_BUCKET" default:"test"`
	AwsAccessKeyID      string        `envconfig:"AWS_ACCESS_KEY_ID" default:"accessKey1"`
	AwsSecretAccessKey  string        `envconfig:"AWS_SECRET_ACCESS_KEY" default:"verySecretKey1"`
	JobCPULimit         string        `envconfig:"JOB_CPU_LIMIT" default:"500m"`
	JobMemoryLimit      string        `envconfig:"JOB_MEMORY_LIMIT" default:"1000Mi"`
	JobCPURequests      string        `envconfig:"JOB_CPU_REQUESTS" default:"300m"`
	JobMemoryRequests   string        `envconfig:"JOB_MEMORY_REQUESTS" default:"400Mi"`
	KubeconfigGenerator string        `envconfig:"KUBECONFIG_GENERATE_IMAGE" default:"quay.io/ocpmetal/ignition-manifests-and-kubeconfig-generate:stable"` // TODO: update the latest once the repository has git workflow
}

func New(log logrus.FieldLogger, kube client.Client, cfg Config) *kubeJob {
	k := &kubeJob{
		Config: cfg,
		log:    log,
		kube:   kube,
	}

	if cfg.ImageBuilderCmd != "" {
		k.imageBuildCmd = strings.Split(cfg.ImageBuilderCmd, " ")
	}

	return k
}

type kubeJob struct {
	Config
	log           logrus.FieldLogger
	kube          client.Client
	imageBuildCmd []string
}

func (k *kubeJob) Create(ctx context.Context, obj runtime.Object, opts ...client.CreateOption) error {
	return k.kube.Create(ctx, obj, opts...)
}

func (k *kubeJob) getJob(ctx context.Context, job *batch.Job, name, namespace string) error {
	retry := func(f func() error) error {
		var err error
		for i := k.RetryAttempts; i > 0; i-- {
			err = f()
			if err == nil {
				return nil
			} else if apierrors.IsNotFound(err) {
				return err
			}
			time.Sleep(k.RetryInterval)
		}
		return err
	}
	//using retry for get job api because sometimes k8s (minikube) api service is not reachable
	if err := retry(func() error {
		return k.kube.Get(ctx, client.ObjectKey{
			Namespace: namespace,
			Name:      name,
		}, job)
	}); err != nil {
		return err
	}
	return nil
}

// Monitor k8s job
func (k *kubeJob) Monitor(ctx context.Context, name, namespace string) error {
	log := logutil.FromContext(ctx, k.log)
	var job batch.Job

	if err := k.getJob(ctx, &job, name, namespace); err != nil {
		return errors.Wrapf(err, "failed to get job <%s>", name)
	}

	for job.Status.Succeeded == 0 && job.Status.Failed < swag.Int32Value(job.Spec.BackoffLimit)+1 {
		time.Sleep(k.MonitorLoopInterval)
		if err := k.getJob(ctx, &job, name, namespace); err != nil {
			return errors.Wrapf(err, "failed to get job <%s>", name)
		}
	}

	if job.Status.Failed >= swag.Int32Value(job.Spec.BackoffLimit)+1 {
		log.Errorf("Job <%s> failed %d times", name, job.Status.Failed)
		return errors.Errorf("Job <%s> failed <%d> times", name, job.Status.Failed)
	}

	// not deleting a job if it failed
	if err := k.kube.Delete(context.Background(), &job); err != nil {
		log.WithError(err).Errorf("Failed to delete job <%s>", name)
	}

	log.Infof("Job <%s> completed successfully", name)
	return nil
}

// Delete k8s job
func (k *kubeJob) Delete(ctx context.Context, name, namespace string) error {
	log := logutil.FromContext(ctx, k.log)
	var job batch.Job

	if err := k.getJob(ctx, &job, name, namespace); err != nil {
		if apierrors.IsNotFound(err) {
			log.Infof("Job <%s> was not found for deletion, probably already completed", name)
			return nil
		}
		log.WithError(err).Errorf("Failed to get job <%s> for deletion", name)
		return errors.Wrapf(err, "failed to get job <%s>", name)
	}

	// not deleting a job if it failed
	if job.Status.Failed >= swag.Int32Value(job.Spec.BackoffLimit)+1 {
		log.Infof("Job <%s> was found already failed", name)
		return nil
	}

	dp := meta.DeletePropagationForeground
	gp := int64(0)
	log.Infof("Sending request to delete job <%s>", name)
	if err := k.kube.Delete(ctx, &job, client.PropagationPolicy(dp), client.GracePeriodSeconds(gp)); err != nil {
		log.WithError(err).Errorf("Failed to delete job <%s>", name)
	}

	// delete is async, wait for the job to not be found
	if err := k.Monitor(ctx, name, namespace); err != nil {
		if !apierrors.IsNotFound(err) {
			log.WithError(err).Errorf("Failed to delete job <%s>", name)
		}
	}
	log.Infof("Completed deletion of job <%s>", name)
	return nil
}

func getQuantity(s string) resource.Quantity {
	reply, _ := resource.ParseQuantity(s)
	return reply
}

// create discovery image generation job, return job name and error
func (k *kubeJob) createImageJob(jobName, imgName, ignitionConfig string) *batch.Job {
	return &batch.Job{
		TypeMeta: meta.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: meta.ObjectMeta{
			Name:      jobName,
			Namespace: k.Config.Namespace,
		},
		Spec: batch.JobSpec{
			BackoffLimit: swag.Int32(2),
			Template: core.PodTemplateSpec{
				ObjectMeta: meta.ObjectMeta{
					Name:      jobName,
					Namespace: k.Config.Namespace,
				},
				Spec: core.PodSpec{
					Containers: []core.Container{
						{
							Resources: core.ResourceRequirements{
								Limits: core.ResourceList{
									"cpu":    getQuantity(k.Config.JobCPULimit),
									"memory": getQuantity(k.Config.JobMemoryLimit),
								},
								Requests: core.ResourceList{
									"cpu":    getQuantity(k.Config.JobCPURequests),
									"memory": getQuantity(k.Config.JobMemoryRequests),
								},
							},
							Name:            "image-creator",
							Image:           k.Config.ImageBuilder,
							Command:         k.imageBuildCmd,
							ImagePullPolicy: "IfNotPresent",
							Env: []core.EnvVar{
								{
									Name:  "S3_ENDPOINT_URL",
									Value: k.Config.S3EndpointURL,
								},
								{
									Name:  "IGNITION_CONFIG",
									Value: ignitionConfig,
								},
								{
									Name:  "IMAGE_NAME",
									Value: imgName,
								},
								{
									Name:  "S3_BUCKET",
									Value: k.Config.S3Bucket,
								},
								{
									Name:  "aws_access_key_id",
									Value: k.Config.AwsAccessKeyID,
								},
								{
									Name:  "aws_secret_access_key",
									Value: k.Config.AwsSecretAccessKey,
								},
							},
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}
}

// creates iso
func (k *kubeJob) GenerateISO(ctx context.Context, cluster common.Cluster, jobName string, imageName string, ignitionConfig string) *installer.GenerateClusterISOInternalServerError {
	log := logutil.FromContext(ctx, k.log)
	if cluster.ID != nil {
		previousCreatedAt := time.Time(cluster.ImageInfo.CreatedAt)
		// Kill the previous job in case it's still running
		prevJobName := fmt.Sprintf("createimage-%s-%s", cluster.ID, previousCreatedAt.Format("20060102150405"))
		log.Info("Attempting to delete job %s", prevJobName)
		if err := k.Delete(ctx, prevJobName, k.Config.Namespace); err != nil {
			log.WithError(err).Errorf("failed to kill previous job in cluster %s", cluster.ID)
			return installer.NewGenerateClusterISOInternalServerError().
				WithPayload(common.GenerateError(http.StatusInternalServerError, err))
		}
		log.Info("Finished attempting to delete job %s", prevJobName)
	}

	// This job name is exactly 63 characters which is the maximum for a job - be careful if modifying
	log.Infof("Creating job %s", jobName)
	if err := k.Create(ctx, k.createImageJob(jobName, imageName, ignitionConfig)); err != nil {
		log.WithError(err).Error("failed to create image job")
		return installer.NewGenerateClusterISOInternalServerError().
			WithPayload(common.GenerateError(http.StatusInternalServerError, err))
	}

	if err := k.Monitor(ctx, jobName, k.Config.Namespace); err != nil {
		log.WithError(err).Error("image creation failed")
		return installer.NewGenerateClusterISOInternalServerError().
			WithPayload(common.GenerateError(http.StatusInternalServerError, err))
	}
	return nil
}

func (k *kubeJob) createKubeconfigJob(cluster *common.Cluster, jobName string, cfg []byte) *batch.Job {
	id := cluster.ID
	kubeConfigGeneratorImage := k.Config.KubeconfigGenerator
	return &batch.Job{
		TypeMeta: meta.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: meta.ObjectMeta{
			Name:      jobName,
			Namespace: k.Config.Namespace,
		},
		Spec: batch.JobSpec{
			BackoffLimit: swag.Int32(2),
			Template: core.PodTemplateSpec{
				ObjectMeta: meta.ObjectMeta{
					Name:      jobName,
					Namespace: k.Config.Namespace,
				},
				Spec: core.PodSpec{
					Containers: []core.Container{
						{
							Name:            kubeconfigPrefix,
							Image:           kubeConfigGeneratorImage,
							Command:         k.imageBuildCmd,
							ImagePullPolicy: "IfNotPresent",
							Env: []core.EnvVar{
								{
									Name:  "S3_ENDPOINT_URL",
									Value: k.Config.S3EndpointURL,
								},
								{
									Name:  "INSTALLER_CONFIG",
									Value: string(cfg),
								},
								{
									Name:  "IMAGE_NAME",
									Value: jobName,
								},
								{
									Name:  "S3_BUCKET",
									Value: k.Config.S3Bucket,
								},
								{
									Name:  "CLUSTER_ID",
									Value: id.String(),
								},
								{
									Name:  "OPENSHIFT_INSTALL_RELEASE_IMAGE_OVERRIDE",
									Value: OverrideImageName, //TODO: change this to match the cluster openshift version
								},
								{
									Name:  "aws_access_key_id",
									Value: k.Config.AwsAccessKeyID,
								},
								{
									Name:  "aws_secret_access_key",
									Value: k.Config.AwsSecretAccessKey,
								},
							},
							Resources: core.ResourceRequirements{
								Limits: core.ResourceList{
									"cpu":    getQuantity(k.Config.JobCPULimit),
									"memory": getQuantity(k.Config.JobMemoryLimit),
								},
								Requests: core.ResourceList{
									"cpu":    getQuantity(k.Config.JobCPURequests),
									"memory": getQuantity(k.Config.JobMemoryRequests),
								},
							},
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}
}

// creates install config
func (k *kubeJob) GenerateInstallConfig(ctx context.Context, cluster common.Cluster, cfg []byte) error {
	log := logutil.FromContext(ctx, k.log)
	jobName := fmt.Sprintf("%s-%s-%s", kubeconfigPrefix, cluster.ID.String(), uuid.New().String())[:63]
	if err := k.Create(ctx, k.createKubeconfigJob(&cluster, jobName, cfg)); err != nil {
		log.WithError(err).Errorf("Failed to create kubeconfig generation job %s for cluster %s", jobName, cluster.ID)
		return errors.Wrapf(err, "Failed to create kubeconfig generation job %s for cluster %s", jobName, cluster.ID)
	}

	if err := k.Monitor(ctx, jobName, k.Config.Namespace); err != nil {
		log.WithError(err).Errorf("Generating kubeconfig files %s failed for cluster %s", jobName, cluster.ID)
		return errors.Wrapf(err, "Generating kubeconfig files %s failed for cluster %s", jobName, cluster.ID)
	}
	return nil
}
