package types

import (
	"github.com/fitan/magic/services/types"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
)

type K8s interface {
	CreateApp(request types.CreateAppRequest) (err error)
	GetApps(keys []types.K8sKey) (res *v1beta1.ApplicationList, err error)
	GetApp(key types.K8sKey) (res *v1beta1.Application, err error)
	GetPodsByDeploymentKey(key types.K8sKey) (pods *v1.PodList, err error)
	GetPodByKey(key types.K8sKey) (res *v1.Pod, err error)
	CreateConfMap(key types.K8sKey, data map[string]string) (res *v1.ConfigMap, err error)
	GetConfigMapByKey(key types.K8sKey) (res *v1.ConfigMap, err error)
	GetConfigMapsByDeploymentKey(key types.K8sKey) (res *v1.ConfigMapList, err error)
	GetDeployment(key types.K8sKey) (res *v12.Deployment, err error)
}
