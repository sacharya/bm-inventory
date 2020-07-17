package installcfg

import (
	"fmt"
	"net"

	"github.com/filanov/bm-inventory/internal/common"
	"github.com/filanov/bm-inventory/models"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type bmc struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type host struct {
	Name            string `yaml:"name"`
	Role            string `yaml:"role"`
	Bmc             bmc    `yaml:"bmc"`
	BootMACAddress  string `yaml:"bootMACAddress"`
	BootMode        string `yaml:"bootMode"`
	HardwareProfile string `yaml:"hardwareProfile"`
}

type baremetal struct {
	ProvisioningNetworkInterface string `yaml:"provisioningNetworkInterface"`
	APIVIP                       string `yaml:"apiVIP"`
	IngressVIP                   string `yaml:"ingressVIP"`
	DNSVIP                       string `yaml:"dnsVIP"`
	Hosts                        []host `yaml:"hosts"`
}

type platform struct {
	Baremetal baremetal `yaml:"baremetal"`
}

type InstallerConfigBaremetal struct {
	APIVersion string `yaml:"apiVersion"`
	BaseDomain string `yaml:"baseDomain"`
	Networking struct {
		NetworkType    string `yaml:"networkType"`
		ClusterNetwork []struct {
			Cidr       string `yaml:"cidr"`
			HostPrefix int    `yaml:"hostPrefix"`
		} `yaml:"clusterNetwork"`
		MachineNetwork []struct {
			Cidr string `yaml:"cidr"`
		} `yaml:"machineNetwork"`
		ServiceNetwork []string `yaml:"serviceNetwork"`
	} `yaml:"networking"`
	Metadata struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Compute []struct {
		Name     string `yaml:"name"`
		Replicas int    `yaml:"replicas"`
	} `yaml:"compute"`
	ControlPlane struct {
		Name     string `yaml:"name"`
		Replicas int    `yaml:"replicas"`
	} `yaml:"controlPlane"`
	Platform            platform `yaml:"platform"`
	PullSecret          string   `yaml:"pullSecret"`
	SSHKey              string   `yaml:"sshKey"`
	ImageContentSources []struct {
		Mirrors []string `yaml:"mirrors"`
		Source  string   `yaml:"source"`
	} `yaml:"imageContentSources"`
	AdditionalTrustBundle string `yaml:"additionalTrustBundle"`
}

func countHostsByRole(cluster *common.Cluster, role models.HostRole) int {
	var count int
	for _, host := range cluster.Hosts {
		if host.Role == role {
			count += 1
		}
	}
	return count
}

func getBasicInstallConfig(cluster *common.Cluster) *InstallerConfigBaremetal {
	return &InstallerConfigBaremetal{
		APIVersion: "v1",
		BaseDomain: cluster.BaseDNSDomain,
		Networking: struct {
			NetworkType    string `yaml:"networkType"`
			ClusterNetwork []struct {
				Cidr       string `yaml:"cidr"`
				HostPrefix int    `yaml:"hostPrefix"`
			} `yaml:"clusterNetwork"`
			MachineNetwork []struct {
				Cidr string `yaml:"cidr"`
			} `yaml:"machineNetwork"`
			ServiceNetwork []string `yaml:"serviceNetwork"`
		}{
			NetworkType: "OpenShiftSDN",
			ClusterNetwork: []struct {
				Cidr       string `yaml:"cidr"`
				HostPrefix int    `yaml:"hostPrefix"`
			}{
				{Cidr: cluster.ClusterNetworkCidr, HostPrefix: int(cluster.ClusterNetworkHostPrefix)},
			},
			MachineNetwork: []struct {
				Cidr string `yaml:"cidr"`
			}{
				{Cidr: cluster.MachineNetworkCidr},
			},
			ServiceNetwork: []string{cluster.ServiceNetworkCidr},
		},
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: cluster.Name,
		},
		Compute: []struct {
			Name     string `yaml:"name"`
			Replicas int    `yaml:"replicas"`
		}{
			{
				Name:     string(models.HostRoleWorker),
				Replicas: countHostsByRole(cluster, models.HostRoleWorker),
			},
		},
		ControlPlane: struct {
			Name     string `yaml:"name"`
			Replicas int    `yaml:"replicas"`
		}{
			Name:     string(models.HostRoleMaster),
			Replicas: countHostsByRole(cluster, models.HostRoleMaster),
		},
		PullSecret: cluster.PullSecret,
		SSHKey:     cluster.SSHPublicKey,
		ImageContentSources: []struct {
			Mirrors []string `yaml:"mirrors"`
			Source  string   `yaml:"source"`
		}{
			{
				Mirrors: []string{
					"local.registry:5000/ocp4",
				},
				Source: "quay.io/openshift-release-dev/ocp-release",
			},
			{
				Mirrors: []string{
					"local.registry:5000/ocp4",
				},
				Source: "quay.io/openshift-release-dev/ocp-v4.0-art-dev",
			},
		},
		AdditionalTrustBundle: "-----BEGIN CERTIFICATE-----\nMIIFzzCCA7egAwIBAgIUQH6Eb1BL3Xi4tLKfeXvAKQrqm8QwDQYJKoZIhvcNAQEL\nBQAwdzELMAkGA1UEBhMCVVMxFjAUBgNVBAgMDU5vcnRoQ2Fyb2xpbmExEDAOBgNV\nBAcMB1JhbGVpZ2gxDzANBgNVBAoMBlJlZEhhdDEUMBIGA1UECwwLRW5naW5lZXJp\nbmcxFzAVBgNVBAMMDmxvY2FsLnJlZ2lzdHJ5MB4XDTIwMDcwODE5MzEyMloXDTMw\nMDcwNjE5MzEyMlowdzELMAkGA1UEBhMCVVMxFjAUBgNVBAgMDU5vcnRoQ2Fyb2xp\nbmExEDAOBgNVBAcMB1JhbGVpZ2gxDzANBgNVBAoMBlJlZEhhdDEUMBIGA1UECwwL\nRW5naW5lZXJpbmcxFzAVBgNVBAMMDmxvY2FsLnJlZ2lzdHJ5MIICIjANBgkqhkiG\n9w0BAQEFAAOCAg8AMIICCgKCAgEAz0/rG+I4FIeGFwnoSNONyjKCF/SCknYlaM0C\n8V56BlHKBRt0pp7jLPvlOOuCk0Db5cwtlYaDpWy0w0iLxWapOClLOS+UDRTAIMTA\nDBsEgww38xLPGpbP4s6yw5SlEuetHPcpIrKckGTeIqapWid51yBf4oF1UsPmaXcc\nAfKj+6+wqwu40dER0G43WTl/CncjsFQJvKQAdu6m0h0Ldqy394moYIBSVsfkDSSq\neP/z6LXGh7M7jF5s5s6BPeGPzuRrgc3IBNJWcUjKFrFXi7u5OmdfG+pg2M3qYwIo\nYHa7sAM8MkRE+Kw6mnzfv0j9V6k2vYu2/ZsVtSwaMklaOheGU6ERTvAZmY7BSwAW\nWMDa9g8XbGJ8dvdoFnFFTHs2YgPG99K2ZUVmxunQaWj/lH3uUGS10wsSmeu5RinA\nBXUY4I+mQpjUg8bqX2V2MitMxwgSYZSX2B3wqFdlTenHLoXiX9aGiSio6KFF1KVQ\nTYJ4gn7LG+PmXgvhtM5+M/TPNMXuGoIiYqTyOv4hLQKRrSIl1MS/1rSjXkGGpYkG\nAImwcIcG74lMtD65AgeD2o8hxMr6uDmI34P8yFuuwn1UJ0+f7Y1eXc/JN/ncgagv\n7lVOzk+E+WyZQrDzdRW201tKBEQVKm/Vxrv4NxbM3dh1SPpRjDWXBS83PbUM86HW\nB7LqSesCAwEAAaNTMFEwHQYDVR0OBBYEFGaknke5Vyb16D8S3b08GqsJooeUMB8G\nA1UdIwQYMBaAFGaknke5Vyb16D8S3b08GqsJooeUMA8GA1UdEwEB/wQFMAMBAf8w\nDQYJKoZIhvcNAQELBQADggIBAGMFNOkVqKMm5YzHQUCwVoByPgn5cnmBQIYe5W6k\n6rBP9FuerBTJF1Kas98IYlH9ctVlpTDFAUuU1zUnOoBGIHbgYIs6htW56ghDTI27\nUCDpAwhv889NvHpWK147z66JHLpTiDYb+34K9sGbuNQDy/YSWLBXIP4uYYChEnnM\nhyjnxO84bphzDzVC57Qj0YLrQaXUX+P1qfNXGS7oIMZH12aAyHcCP/qerTbYAVPy\nUmGziGaF4S/Sg1c28yKL5eGE1fwHDLYqiedGKZwMRgVt4b3m7GoVjVmkhCCnKn4O\nRuoAGETuRHpI/6S0e4UaUDMi8D0F+wNP/AAikTSB2339EVA57MXeuqAByKjbyiR2\n3d8mvgiIm2CMTxomNpc0Fg+90tLdsQpwXhMIF7q3JVq55ri8RjhdyYYtWq03omx2\nXNVLb8igICsHeVFgBaJ54GpAwgXsP+fcPCXzn2I4Gj0zmP10z942wmIWwgGbJAI+\nL8kD5lyFxCOuZYUn3tZgSa4S6GR/MWJPMZwht2oUQyDLs0OJ7HEQKRN/OOkJq3An\nANtGSADHldAqbZq5KfaPcdKXERAxrpmoKhazBV+joDlF138Uke6jnUZ2ja83Zz/9\nnG52KdhuWNWCZgh6kgw/1ULYmoog6kiBLGF5nk4GBww8GaKVw7c3mU71BHgNufWW\neOod\n-----END CERTIFICATE-----",
	}
}

// [TODO] - remove once we decide to use specific values from the hosts of the cluster
func getDummyMAC(log logrus.FieldLogger, dummyMAC string, count int) (string, error) {
	hwMac, err := net.ParseMAC(dummyMAC)
	if err != nil {
		log.Warn("Failed to parse dummyMac")
		return "", err
	}
	hwMac[len(hwMac)-1] = hwMac[len(hwMac)-1] + byte(count)
	return hwMac.String(), nil
}

func setBMPlatformInstallconfig(log logrus.FieldLogger, cluster *common.Cluster, cfg *InstallerConfigBaremetal) error {
	// set hosts
	numMasters := countHostsByRole(cluster, models.HostRoleMaster)
	numWorkers := countHostsByRole(cluster, models.HostRoleWorker)
	masterCount := 0
	workerCount := 0
	hosts := make([]host, numWorkers+numMasters)

	// dummy MAC and port, once we start using real BMH, those values should be set from cluster
	dummyMAC := "00:aa:39:b3:51:10"
	dummyPort := 6230

	for i := range hosts {
		log.Infof("Setting master, host %d, master count %d", i, masterCount)
		if i >= numMasters {
			hosts[i].Name = fmt.Sprintf("openshift-worker-%d", workerCount)
			hosts[i].Role = string(models.HostRoleWorker)
			workerCount += 1
		} else {
			hosts[i].Name = fmt.Sprintf("openshift-master-%d", masterCount)
			hosts[i].Role = string(models.HostRoleMaster)
			masterCount += 1
		}
		hosts[i].Bmc = bmc{
			Address:  fmt.Sprintf("ipmi://192.168.111.1:%d", dummyPort+i),
			Username: "admin",
			Password: "rackattack",
		}
		hwMac, err := getDummyMAC(log, dummyMAC, i)
		if err != nil {
			log.Warn("Failed to parse dummyMac")
			return err
		}
		hosts[i].BootMACAddress = hwMac
		hosts[i].BootMode = "UEFI"
		hosts[i].HardwareProfile = "unknown"
	}
	cfg.Platform = platform{
		Baremetal: baremetal{
			ProvisioningNetworkInterface: "ethh0",
			APIVIP:                       cluster.APIVip,
			IngressVIP:                   cluster.IngressVip,
			DNSVIP:                       cluster.APIVip,
			Hosts:                        hosts,
		},
	}
	return nil
}

func GetInstallConfig(log logrus.FieldLogger, cluster *common.Cluster) ([]byte, error) {
	cfg := getBasicInstallConfig(cluster)
	err := setBMPlatformInstallconfig(log, cluster, cfg)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(*cfg)
}
