package k8s

import (
	"fmt"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/fitan/magic/pkg/types"
	types2 "github.com/fitan/magic/services/types"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"io"
	"k8s.io/api/core/v1"
	"strings"
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

type WatchPodLogsIn struct {
	Uri struct {
		types2.K8sKey
		PodName       string `json:"podName" uri:"podName"`
		ContainerName string `json:"containerName" uri:"containerName"`
	} `json:"uri"`
}

// @Description Get pod logs
// @GenApi /k8s/:namespace/app/:name/pod/:podName/container/:containerName/logs [get]
func WatchPodLogs(core *types.Core, in *WatchPodLogsIn) (string, error) {
	c := make(chan string, 0)
	go core.GetServices().K8s().WatchPodLogs(in.Uri.K8sKey, in.Uri.PodName, in.Uri.ContainerName, c)
	core.GinX.GinCtx().Stream(
		func(w io.Writer) bool {
			if msg, ok := <-c; ok {
				core.GinX.GinCtx().SSEvent("log", strings.Replace(msg, "\n", "", -1))
				return true
			}
			return false
		})
	return "", nil
}

type DownloadPodLogsIn struct {
	Uri struct {
		types2.K8sKey
		PodName       string `json:"podName" uri:"podName"`
		ContainerName string `json:"containerName" uri:"containerName"`
	} `json:"uri"`
	Query struct {
		TailLines *int64 `json:"tailLines" form:"tailLines"`
	} `json:"query"`
}

// @Description download pod logs
// @GenApi /k8s/:namespace/app/:name/pod/:podName/container/:containerName/logs/download [get]
func DownloadPodLogs(core *types.Core, in *DownloadPodLogsIn) (string, error) {
	log, err := core.GetServices().K8s().GetPodLogs(in.Uri.K8sKey, in.Uri.PodName, in.Uri.ContainerName, in.Query.TailLines)
	if err != nil {
		return log, err
	}
	core.GetGinX().GinCtx().Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", in.Uri.PodName+".log"))
	core.GetGinX().GinCtx().Data(200, "Content-Type: text/xml", []byte(log))
	return "", ginx.SkipWrapError
}
