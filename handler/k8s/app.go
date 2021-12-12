package k8s

import (
	"github.com/fitan/magic/pkg/types"
	types2 "github.com/fitan/magic/services/types"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
)

type SpaceName struct {
	Namespace string `json:"namespace" uri:"namespace"`
	Name      string `json:"name" uri:"name"`
}

type GetAppIn struct {
	Uri SpaceName
}

// @Description 获取app
// @GenApi /k8s/:namespace/app/:name [get]
func GetApp(core *types.Core, in *GetAppIn) (*v1beta1.Application, error) {
	return core.Services.K8s().GetApp(types2.K8sKey{
		Namespace: in.Uri.Namespace,
		Name:      in.Uri.Name,
	})
}
