package k8s

import (
	types2 "github.com/fitan/magic/services/types"
	"net/http"

	"github.com/fitan/magic/handler/k8s"
	"github.com/fitan/magic/pkg/types"
)

type GetAppTransfer struct {
}

func (t *GetAppTransfer) Method() string {
	return http.MethodGet
}

func (t *GetAppTransfer) Url() string {
	return "/k8s/:namespace/app/:name"
}

func (t *GetAppTransfer) Binder() types.GinXBinder {
	return new(GetAppBinder)
}

type GetAppBinder struct {
	val *k8s.GetAppIn
}

func (b *GetAppBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.GetAppIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param namespace path string true " "
// @Param name path string true " "
// @Success 200 {object} ginx.GinXResult{data=v1beta1.Application}
// @Description 获取app
// @Router /k8s/:namespace/app/:name [get]
func (b *GetAppBinder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.GetApp(core, b.val)
}

type SwagCreateWorkerBody types2.Worker

type CreateWorkerTransfer struct {
}

func (t *CreateWorkerTransfer) Method() string {
	return http.MethodPost
}

func (t *CreateWorkerTransfer) Url() string {
	return "/k8s/:namespace/app/:name"
}

func (t *CreateWorkerTransfer) Binder() types.GinXBinder {
	return new(CreateWorkerBinder)
}

type CreateWorkerBinder struct {
	val *k8s.CreateWorkerIn
}

func (b *CreateWorkerBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.CreateWorkerIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param body body SwagCreateWorkerBody true " "
// @Param namespace path string true " "
// @Param name path string true " "
// @Success 200 {object} ginx.GinXResult{data=bool}
// @Description 创建worker
// @Router /k8s/:namespace/app/:name [post]
func (b *CreateWorkerBinder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.CreateWorker(core, b.val)
}

type GetPodsTransfer struct {
}

func (t *GetPodsTransfer) Method() string {
	return http.MethodGet
}

func (t *GetPodsTransfer) Url() string {
	return "/k8s/:namespace/app/:name/pod"
}

func (t *GetPodsTransfer) Binder() types.GinXBinder {
	return new(GetPodsBinder)
}

type GetPodsBinder struct {
	val *k8s.GetPodsIn
}

func (b *GetPodsBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.GetPodsIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param namespace path string true " "
// @Param name path string true " "
// @Success 200 {object} ginx.GinXResult{data=v1.PodList}
// @Description Get Pods
// @Router /k8s/:namespace/app/:name/pod [get]
func (b *GetPodsBinder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.GetPods(core, b.val)
}

type WatchPodLogsTransfer struct {
}

func (t *WatchPodLogsTransfer) Method() string {
	return http.MethodGet
}

func (t *WatchPodLogsTransfer) Url() string {
	return "/k8s/:namespace/app/:name/pod/:podName/container/:containerName/logs"
}

func (t *WatchPodLogsTransfer) Binder() types.GinXBinder {
	return new(WatchPodLogsBinder)
}

type WatchPodLogsBinder struct {
	val *k8s.WatchPodLogsIn
}

func (b *WatchPodLogsBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.WatchPodLogsIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param namespace path string true " "
// @Param name path string true " "
// @Param podName path string true " "
// @Param containerName path string true " "
// @Success 200 {object} ginx.GinXResult{data=string}
// @Description Get pod logs
// @Router /k8s/:namespace/app/:name/pod/:podName/container/:containerName/logs [get]
func (b *WatchPodLogsBinder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.WatchPodLogs(core, b.val)
}

type SwagDownloadPodLogsQuery struct {
	TailLines int64 `json:"tailLines" form:"tailLines"`
}

type DownloadPodLogsTransfer struct {
}

func (t *DownloadPodLogsTransfer) Method() string {
	return http.MethodGet
}

func (t *DownloadPodLogsTransfer) Url() string {
	return "/k8s/:namespace/app/:name/pod/:podName/container/:containerName/logs/download"
}

func (t *DownloadPodLogsTransfer) Binder() types.GinXBinder {
	return new(DownloadPodLogsBinder)
}

type DownloadPodLogsBinder struct {
	val *k8s.DownloadPodLogsIn
}

func (b *DownloadPodLogsBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.DownloadPodLogsIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	err = core.GinX.GinCtx().ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param query query SwagDownloadPodLogsQuery false " "
// @Param namespace path string true " "
// @Param name path string true " "
// @Param podName path string true " "
// @Param containerName path string true " "
// @Success 200 {object} ginx.GinXResult{data=string}
// @Description download pod logs
// @Router /k8s/:namespace/app/:name/pod/:podName/container/:containerName/logs/download [get]
func (b *DownloadPodLogsBinder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.DownloadPodLogs(core, b.val)
}
