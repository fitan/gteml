package types

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)

type Labels map[string]string

func (t *Labels) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "labels",
		Properties: util.Object2RawExtension(*t),
	}
}

type Annotations map[string]string

func (t *Annotations) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "annotations",
		Properties: util.Object2RawExtension(*t),
	}
}

type Sidecar struct {
	Name    string    `json:"name"`
	Image   string    `json:"image"`
	Cmd     *[]string `json:"cmd"`
	Args    *[]string `json:"args"`
	Volumes *[]struct {
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"volumes"`
}

func (s *Sidecar) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "sidecar",
		Properties: util.Object2RawExtension(s),
	}
}

type Scaler struct {
	Replicas int `json:"replicas"`
}

func (s *Scaler) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "scaler",
		Properties: util.Object2RawExtension(s),
	}
}

type InitContainer struct {
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Cmd           *[]string `json:"cmd"`
	Args          *[]string `json:"args"`
	MountName     string    `json:"mountName"`
	AppMountPath  string    `json:"appMountPath"`
	InitMountPath string    `json:"initMountPath"`
}

func (i *InitContainer) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "init-container",
		Properties: util.Object2RawExtension(i),
	}
}

type Expose struct {
	Port []int  `json:"port"`
	Type string `json:"type"`
}

func (e *Expose) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "expose",
		Properties: util.Object2RawExtension(e),
	}
}

type ConfigMap struct {
	Volumes []struct {
		Name      string             `json:"name"`
		MountPath string             `json:"mountPath"`
		readOnly  bool               `json:"readOnly"`
		Data      *map[string]string `json:"data"`
	} `json:"volumes"`
}

func (c *ConfigMap) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "configmap",
		Properties: util.Object2RawExtension(c),
	}
}

type Pvc struct {
	ClaimName        string   `json:"claimName"`
	VolumeMode       string   `json:"volumeMode"`
	VolumeName       *string  `json:"volumeName"`
	AccessModes      []string `json:"accessModes"`
	StorageClassName *string  `json:"storageClassName"`
	Resources        struct {
		Requests string `json:"requests"`
		Limits   string `json:"limits"`
	} `json:"resources"`

	VolumesToMount []struct {
		Name       string `json:"name"`
		DevicePath string `json:"devicePath"`
		MountPath  string `json:"mountPath"`
	} `json:"volumesToMount"`
}

func (p *Pvc) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "pvc",
		Properties: util.Object2RawExtension(p),
	}
}

type MyPorts struct {
	Ports []struct {
		ContainerPort int    `json:"containerPort"`
		Ptotocol      string `json:"ptotocol"`
	} `json:"ports"`
}

func (p *MyPorts) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "my-ports",
		Properties: util.Object2RawExtension(p),
	}
}

type MyEnv struct {
}

func (m *MyEnv) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       "my-env",
		Properties: util.Object2RawExtension(m),
	}
}
