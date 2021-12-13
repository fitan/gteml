package types

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)

type K8sKey struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}

type CreatePvOption struct {
}

type CreateAppRequest struct {
	Metadata struct {
		Name      string `json:"name"`
		NameSpace string `json:"nameSpace"`
	} `json:"metadata"`
	Components []struct {
		Name       string `json:"name"`
		Type       string `json:"type"`
		Properties struct {
			Image           string   `json:"image"`
			Cmd             []string `json:"cmd"`
			Port            int      `json:"port"`
			Cpu             string   `json:"cpu"`
			Mem             string   `json:"mem"`
			ImagePullPolicy string   `json:"imagePullPolicy"`
			Volumes         []struct {
				Name      string `json:"name"`
				MountPath string `json:"mountPath"`
				Type      string `json:"type"`
			} `json:"volumes"`
			Env []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"env"`
			LivenessProbe struct {
				Exec    []string `json:"exec"`
				HttpGet struct {
					Path        string
					Port        string
					httpHeaders []struct {
						Name  string
						Value string
					}
				}
				TcpSocket struct {
					Port int
				}
				InitialDelaySeconds int
				PeriodSeconds       int
				TimeoutSeconds      int
				SuccessThreshold    int
				FailureThreshold    int
			} `json:"livenessProbe"`
			ReadinessProbe struct{} `json:"readinessProbe"`
		}
		Traits struct {
			Ingress TraitIngress `json:"ingress"`

			Rollout TraitRollout `json:"rollout"`

			CpuScaler TraitCpuScaler `json:"cpu_scaler"`

			Labels TraitLabels `json:"labels"`

			Annotations TraitAnnotations `json:"annotations"`

			Sidecar TraitSidecar `json:"sidecar"`
		}
	} `json:"components"`
}

type TraitIngress struct {
	Domain string         `json:"domain"`
	Http   map[string]int `json:"http"`
}

func (t *TraitIngress) ToTrait() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type: "ingress",
		Properties: util.Object2RawExtension(map[string]interface{}{
			"domain": t.Domain,
			"http":   t.Http,
		}),
	}
}

type TraitRollout struct {
	TargetSize     int   `json:"targetSize"`
	RolloutBatches []int `json:"rolloutBatches"`
}

func (t *TraitRollout) ToTrait() common.ApplicationTrait {
	rolloutBatches := make([]map[string]int, 0, len(t.RolloutBatches))
	for _, i := range t.RolloutBatches {
		replicas := i
		rolloutBatches = append(rolloutBatches, map[string]int{"- replicas": replicas})
	}
	return common.ApplicationTrait{
		Type: "rollout",
		Properties: util.Object2RawExtension(map[string]interface{}{
			"targetSize":     t.TargetSize,
			"rolloutBatches": rolloutBatches,
		}),
	}
}

type TraitCpuScaler struct {
	Min        uint `json:"min"`
	Max        uint `json:"max"`
	CpuPercent uint `json:"cpuPercent"`
}

type TraitLabels map[string]string

type TraitAnnotations map[string]string

type TraitServiceBinding map[string]string

type TraitSidecar struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Cmd     []string `json:"cmd"`
	Volumes []struct {
		Name string `json:"name"`
		Path string `json:"path"`
	}
}
