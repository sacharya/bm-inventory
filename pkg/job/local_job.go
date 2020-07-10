package job

import (
	"bytes"
	"context"
	"os"
	"os/exec"

	"github.com/filanov/bm-inventory/internal/common"
	logutil "github.com/filanov/bm-inventory/pkg/log"
	"github.com/filanov/bm-inventory/restapi/operations/installer"
	"github.com/sirupsen/logrus"
)

type localJob struct {
	Config
	log logrus.FieldLogger
}

func NewLocalJob(log logrus.FieldLogger, cfg Config) *localJob {
	return &localJob{
		Config: cfg,
		log:    log,
	}
}

// creates install config
func (j *localJob) GenerateInstallConfig(ctx context.Context, cluster common.Cluster, cfg []byte) error {
	log := logutil.FromContext(ctx, j.log)
	cmd := exec.Command("python", "./data/process-ignition-manifests-and-kubeconfig.py")
	cmd.Env = append(os.Environ(),
		"S3_ENDPOINT_URL="+j.Config.S3EndpointURL,
		"INSTALLER_CONFIG="+string(cfg),
		"IMAGE_NAME="+j.Config.KubeconfigGenerator,
		"S3_BUCKET="+j.Config.S3Bucket,
		"CLUSTER_ID="+cluster.ID.String(),
		"OPENSHIFT_INSTALL_RELEASE_IMAGE_OVERRIDE="+OverrideImageName,
		"aws_access_key_id="+j.Config.AwsAccessKeyID,
		"aws_secret_access_key="+j.Config.AwsSecretAccessKey,
		"WORK_DIR=/data",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	log.Println(cmd.Stdout)
	log.Println(cmd.Env)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (j *localJob) GenerateISO(ctx context.Context, cluster common.Cluster, jobName string, imageName string, ignitionConfig string) *installer.GenerateClusterISOInternalServerError {
	log := logutil.FromContext(ctx, j.log)
	cmd := exec.Command("python", "./data/install_process.py")
	cmd.Env = append(os.Environ(),
		"S3_ENDPOINT_URL="+j.Config.S3EndpointURL,
		"IGNITION_CONFIG="+ignitionConfig,
		"IMAGE_NAME="+imageName,
		"S3_BUCKET="+j.Config.S3Bucket,
		"aws_access_key_id="+j.Config.AwsAccessKeyID,
		"aws_secret_access_key="+j.Config.AwsSecretAccessKey,
		"WORK_DIR=/data",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	log.Println(cmd.Stdout)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return nil
}
