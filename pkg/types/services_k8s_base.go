package types

import (
	"github.com/fitan/magic/services/types"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
)

type K8s interface {
	CreateApp(request types.CreateAppRequest) (err error)
	GetApp(namespace, name string) (app *v1beta1.Application, err error)
	GetPodsByDeploymentName(namespace string, deploymentName string) (pods *v1.PodList, err error)
	GetPod(namespace string, name string) (res *v1.Pod, err error)
	CreateConfMap(namespace string, name string, data map[string]string) (res *v1.ConfigMap, err error)
	GetConfigMap(namespace, name string) (res *v1.ConfigMap, err error)
	GetConfigMapsByDeployment(namespace string, deploymentName string) (res *v1.ConfigMapList, err error)
	GetDeployment(namespace, name string) (res *v12.Deployment, err error)
}
