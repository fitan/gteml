package k8s

import (
	"github.com/fitan/magic/pkg/types"
	types2 "github.com/fitan/magic/services/types"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"k8s.io/api/core/v1"
)

type SpaceName struct {
	Namespace string `json:"namespace" uri:"namespace"`
	Name      string `json:"name" uri:"name"`
}

type GetAppIn struct {
	Uri types2.K8sKey
}

// @Description 获取app
// @GenApi /k8s/:namespace/app/:name [get]
func GetApp(core *types.Core, in *GetAppIn) (*v1beta1.Application, error) {
	return core.Services.K8s().GetApp(in.Uri)
}

type CreateWorkerIn struct {
	Uri  types2.K8sKey `json:"uri"`
	Body types2.Worker `json:"body"`
}

// @Description 创建worker
// @GenApi /k8s/:namespace/app/:name [post]
func CreateWorker(core *types.Core, in *CreateWorkerIn) (bool, error) {
	err := core.Services.K8s().ApplyWorker(&in.Body)
	if err != nil {
		return false, err
	}
	return true, err

}

type GetPodsIn struct {
	Uri types2.K8sKey `json:"uri"`
}

// @Description Get Pods
// @GenApi /k8s/:namespace/app/:name/pod [get]
func GetPods(core *types.Core, in *GetPodsIn) (*v1.PodList, error) {
	return core.Services.K8s().GetPodsByDeploymentKey(in.Uri)
}
