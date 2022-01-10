package k8s

import (
	"net/http"

	"github.com/fitan/magic/handler/k8s"
	"github.com/fitan/magic/pkg/types"
	types2 "github.com/fitan/magic/services/types"
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

type SwagDownloadPodFileQuery struct {
	FilePath string `json:"filePath" form:"filePath"`
}

type DownloadPodFileTransfer struct {
}

func (t *DownloadPodFileTransfer) Method() string {
	return http.MethodGet
}

func (t *DownloadPodFileTransfer) Url() string {
	return "/k8s/:namespace/app/:name/pod/:podName/container/:containerName/file"
}

func (t *DownloadPodFileTransfer) Binder() types.GinXBinder {
	return new(DownloadPodFileBinder)
}

type DownloadPodFileBinder struct {
	val *k8s.DownloadPodFileIn
}

func (b *DownloadPodFileBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.DownloadPodFileIn{}
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
// @Param query query SwagDownloadPodFileQuery false " "
// @Param namespace path string true " "
// @Param name path string true " "
// @Param podName path string true " "
// @Param containerName path string true " "
// @Success 200 {object} ginx.GinXResult{data=string}
// @Description 下载pod里的文件
// @Router /k8s/:namespace/app/:name/pod/:podName/container/:containerName/file [get]
func (b *DownloadPodFileBinder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.DownloadPodFile(core, b.val)
}

type SwagDownloadPodFileV2Query struct {
	FilePath string `json:"filePath" form:"filePath"`
}

type DownloadPodFileV2Transfer struct {
}

func (t *DownloadPodFileV2Transfer) Method() string {
	return http.MethodGet
}

func (t *DownloadPodFileV2Transfer) Url() string {
	return "/k8s/:namespace/app/:name/pod/:podName/container/:containerName/file/v2"
}

func (t *DownloadPodFileV2Transfer) Binder() types.GinXBinder {
	return new(DownloadPodFileV2Binder)
}

type DownloadPodFileV2Binder struct {
	val *k8s.DownloadPodFileIn
}

func (b *DownloadPodFileV2Binder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.DownloadPodFileIn{}
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
// @Param query query SwagDownloadPodFileV2Query false " "
// @Param namespace path string true " "
// @Param name path string true " "
// @Param podName path string true " "
// @Param containerName path string true " "
// @Success 200 {object} ginx.GinXResult{data=int64}
// @Description 下载pod里的文件 V2
// @Router /k8s/:namespace/app/:name/pod/:podName/container/:containerName/file/v2 [get]
func (b *DownloadPodFileV2Binder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.DownloadPodFileV2(core, b.val)
}

type SwagPortforwardQuery struct {
	Ports []string `json:"ports" form:"ports"`
}

type PortforwardTransfer struct {
}

func (t *PortforwardTransfer) Method() string {
	return http.MethodGet
}

func (t *PortforwardTransfer) Url() string {
	return "/k8s/:namespace/app/:name/pod/:podName/portforward"
}

func (t *PortforwardTransfer) Binder() types.GinXBinder {
	return new(PortforwardBinder)
}

type PortforwardBinder struct {
	val *k8s.PortforwardIn
}

func (b *PortforwardBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.PortforwardIn{}
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
// @Param query query SwagPortforwardQuery false " "
// @Param namespace path string true " "
// @Param name path string true " "
// @Param podName path string true " "
// @Success 200 {object} ginx.GinXResult{data=string}
// @Description
// @Router /k8s/:namespace/app/:name/pod/:podName/portforward [get]
func (b *PortforwardBinder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.Portforward(core, b.val)
}
