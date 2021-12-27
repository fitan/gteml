package types

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	appv1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ApplicationKindName = "Application"

type K8sKey struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}

type CreatePvOption struct {
}

type ToTraiter interface {
	ToTrait() common.ApplicationTrait
}

type CreateAppRequest struct {
	Metadata struct {
		Name      string `json:"name"`
		NameSpace string `json:"nameSpace"`
	} `json:"metadata"`
	Components []Component `json:"components"`
}

func (a *CreateAppRequest) ToApplication() *appv1beta1.Application {
	components := make([]common.ApplicationComponent, 0, 0)
	for _, v := range a.Components {
		components = append(components, v.ToApplicationComponent())
	}

	return &appv1beta1.Application{
		TypeMeta: v1.TypeMeta{
			Kind:       "Application",
			APIVersion: "core.oam.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Namespace: a.Metadata.NameSpace,
			Name:      a.Metadata.Name,
		},
		Spec: appv1beta1.ApplicationSpec{
			Components: components,
			Policies: []appv1beta1.AppPolicy{appv1beta1.AppPolicy{
				Name:       "health-policy" + "-" + a.Metadata.Name,
				Type:       "health",
				Properties: util.Object2RawExtension(map[string]interface{}{"probeInterval": 30, "probeTimeout": 10}),
			}},
		},
	}
}

type Component struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Image           string   `json:"image"`
		Cmd             []string `json:"cmd"`
		Port            int      `json:"port"`
		Cpu             string   `json:"cpu"`
		Mem             string   `json:"mem"`
		ImagePullPolicy *string  `json:"imagePullPolicy"`
		Volumes         *[]struct {
			Name      string `json:"name"`
			MountPath string `json:"mountPath"`
			Type      string `json:"type"`
		} `json:"volumes"`
		Env *[]struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"env"`
		LivenessProbe  *HealthProbe `json:"livenessProbe"`
		ReadinessProbe *HealthProbe `json:"readinessProbe"`
	}
	Traits struct {
		Ingress *TraitIngress `json:"ingress"`

		Rollout *TraitRollout `json:"rollout"`

		//CpuScaler *TraitCpuScaler `json:"cpu_scaler"`

		Labels *Labels `json:"labels"`

		Annotations *Annotations `json:"annotations"`

		Sidecar *Sidecar `json:"sidecar"`
	}
}

type HealthProbe struct {
	Exec    *[]string `json:"exec"`
	HttpGet *struct {
		Path        string `json:"path"`
		Port        string `json:"port"`
		HttpHeaders *[]struct {
			Name  string
			Value string
		} `json:"httpHeaders"`
	} `json:"httpGet"`
	TcpSocket *struct {
		Port int `json:"port"`
	} `json:"tcpSocket"`
	InitialDelaySeconds int `json:"initialDelaySeconds"`
	PeriodSeconds       int `json:"periodSeconds"`
	TimeoutSeconds      int `json:"timeoutSeconds"`
	SuccessThreshold    int `json:"successThreshold"`
	FailureThreshold    int `json:"failureThreshold"`
}

func (c *Component) ToApplicationComponent() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0, 0)
	if c.Traits.Ingress != nil {
		traits = append(traits, c.Traits.Ingress.ToTrait())
	}

	if c.Traits.Annotations != nil {
		traits = append(traits, c.Traits.Annotations.ToTrait())
	}

	if c.Traits.Labels != nil {
		traits = append(traits, c.Traits.Labels.ToTrait())
	}

	if c.Traits.Rollout != nil {
		traits = append(traits, c.Traits.Rollout.ToTrait())
	}

	if c.Traits.Sidecar != nil {
		traits = append(traits, c.Traits.Sidecar.ToTrait())
	}
	return common.ApplicationComponent{
		Name:       c.Name,
		Type:       "webservice",
		Properties: util.Object2RawExtension(c.Properties),
		Traits:     traits,
	}
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

//type TraitCpuScaler struct {
//	Min        uint `json:"min"`
//	Max        uint `json:"max"`
//	CpuPercent uint `json:"cpuPercent"`
//}
//
//func (t *TraitCpuScaler) ToTrait() common.ApplicationTrait {
//	return common.ApplicationTrait{
//		Type: "cpuscaler",
//		Properties: util.Object2RawExtension(map[string]interface{}{
//			"min":        t.Min,
//			"max":        t.Max,
//			"cpuPercent": t.CpuPercent,
//		}),
//	}
//}

type TraitServiceBinding map[string]string

func (t *TraitServiceBinding) ToTrait() common.ApplicationTrait {
	envMap := make(map[string]map[string]string)
	for k, v := range *t {
		value := map[string]string{"secret": v}
		envMap[k] = value
	}
	return common.ApplicationTrait{
		Type: "service-binding",
		Properties: util.Object2RawExtension(map[string]interface{}{
			"envMappings": envMap,
		}),
	}
}
