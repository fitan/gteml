package types

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	appv1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

type Worker struct {
	Metadata  K8sKey          `json:"metadata"`
	Component WorkerComponent `json:"component"`
}

type WorkerComponent struct {
	Properties struct {
		Image            string    `json:"image"`
		ImagePullPolicy  *string   `json:"imagePullPolicy,omitempty"`
		ImagePullSecrets *string   `json:"imagePullSecrets,omitempty"`
		Cmd              *[]string `json:"cmd,omitempty"`
		Env              *[]struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"env,omitempty"`
		Cpu    *string `json:"cpu,omitempty"`
		Memory *string `json:"memory,omitempty"`

		Volumes *[]struct {
			Name      string `json:"name"`
			MountPath string `json:"mountPath"`
			Type      string `json:"type"`
		} `json:"volumes,omitempty"`

		LivenessProbe  *HealthProbe `json:"livenessProbe,omitempty"`
		ReadinessProbe *HealthProbe `json:"readinessProbe,omitempty"`
	} `json:"properties"`

	Traits struct {
		Ingress       *TraitIngress  `json:"ingress,omitempty"`
		Labels        *Labels        `json:"labels,omitempty"`
		Annotations   *Annotations   `json:"annotations,omitempty"`
		Sidecar       *Sidecar       `json:"sidecar,omitempty"`
		Expose        *Expose        `json:"expose,omitempty"`
		InitContainer *InitContainer `json:"initContainer,omitempty"`
		ConfigMap     *ConfigMap     `json:"configMap,omitempty"`
		Pvc           *Pvc           `json:"pvc,omitempty"`
		Scaler        *Scaler        `json:"scaler,omitempty"`
		Ports         *MyPorts       `json:"ports,omitempty"`
		MyEnv         *MyEnv         `json:"myEnv,omitempty"`
		Test          *struct{}      `json:"test,omitempty"`
	} `json:"traits,omitempty"`
}

func (w *Worker) ToWorker() *appv1beta1.Application {
	return &appv1beta1.Application{
		TypeMeta: v1.TypeMeta{
			Kind:       "Application",
			APIVersion: "core.oam.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Namespace: w.Metadata.Namespace,
			Name:      w.Metadata.Name,
		},
		Spec: appv1beta1.ApplicationSpec{
			Components: []common.ApplicationComponent{w.ToComponent()},
			Policies: []appv1beta1.AppPolicy{appv1beta1.AppPolicy{
				Name:       "health-policy" + "-" + w.Metadata.Name,
				Type:       "health",
				Properties: util.Object2RawExtension(map[string]interface{}{"probeInterval": 30, "probeTimeout": 10}),
			}},
		},
	}
}

func (w *Worker) ToComponent() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0, 0)

	v := reflect.ValueOf(w.Component.Traits)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			continue
		}

		toTraiter, ok := v.Field(i).Interface().(ToTraiter)
		if !ok {
			panic(v.Field(i).String() + "Do not implement ToTraiter")
		}

		trait := toTraiter.ToTrait()
		traits = append(traits, trait)

	}

	//traits := make([]common.ApplicationTrait, 0, 0)
	//if w.Component.Traits.ConfigMap != nil {
	//	traits = append(traits, w.Component.Traits.ConfigMap.ToTrait())
	//}
	//
	//if w.Component.Traits.Pvc != nil {
	//	traits = append(traits, w.Component.Traits.Pvc.ToTrait())
	//}
	//
	//if w.Component.Traits.Expose != nil {
	//	traits = append(traits, w.Component.Traits.Expose.ToTrait())
	//}
	//
	//if w.Component.Traits.Labels != nil {
	//	traits = append(traits, w.Component.Traits.Labels.ToTrait())
	//}
	//
	//if w.Component.Traits.Annotations != nil {
	//	traits = append(traits, w.Component.Traits.Annotations.ToTrait())
	//}
	//
	//if w.Component.Traits.InitContainer != nil {
	//	traits = append(traits, w.Component.Traits.InitContainer.ToTrait())
	//}
	//
	//if w.Component.Traits.Sidecar != nil {
	//	traits = append(traits, w.Component.Traits.Sidecar.ToTrait())
	//}
	//
	//if w.Component.Traits.Ingress != nil {
	//	traits = append(traits, w.Component.Traits.Ingress.ToTrait())
	//}
	//
	//if w.Component.Traits.Scaler != nil {
	//	traits = append(traits, w.Component.Traits.Scaler.ToTrait())
	//}
	//
	//if w.Component.Traits.Ports != nil {
	//	traits = append(traits, w.Component.Traits.Ports.ToTrait())
	//}
	//
	//if w.Component.Traits.MyEnv != nil {
	//	traits = append(traits, w.Component.Traits.MyEnv.ToTrait())
	//}

	return common.ApplicationComponent{
		Name:       w.Metadata.Name,
		Type:       "worker",
		Properties: util.Object2RawExtension(w.Component.Properties),
		Traits:     traits,
	}
}
