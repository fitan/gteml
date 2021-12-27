package types

import (
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"k8s.io/apimachinery/pkg/runtime"
)

type Worker struct {
	Component WorkerComponent
	Traiter   struct {
		Ingress     *TraitIngress `json:"ingress"`
		Labels      *Labels       `json:"labels"`
		Annotations *Annotations  `json:"annotations"`
		Sidecar     *Sidecar      `json:"sidecar"`
	}
}

type WorkerComponent struct {
	Image            string   `json:"image"`
	ImagePullPolicy  string   `json:"imagePullPolicy"`
	ImagePullSecrets string   `json:"imagePullSecrets"`
	Cmd              []string `json:"cmd"`
	Env              []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"env"`
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`

	Volumes *[]struct {
		Name      string `json:"name"`
		MountPath string `json:"mountPath"`
		Type      string `json:"type"`
	} `json:"volumes"`

	LivenessProbe  *HealthProbe `json:"livenessProbe"`
	ReadinessProbe *HealthProbe `json:"readinessProbe"`
}

func (w *WorkerComponent) ToProperties() *runtime.RawExtension {
	return util.Object2RawExtension(w)
}
