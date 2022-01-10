package k8s

import (
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter, reg *ginx.GinXHandlerRegister) {

	reg.Register(r, &GetAppTransfer{}, ginx.WithHandlerName("GetApp"))

	reg.Register(r, &CreateWorkerTransfer{}, ginx.WithHandlerName("CreateWorker"))

	reg.Register(r, &GetPodsTransfer{}, ginx.WithHandlerName("GetPods"))

	reg.Register(r, &WatchPodLogsTransfer{}, ginx.WithHandlerName("WatchPodLogs"))

	reg.Register(r, &DownloadPodLogsTransfer{}, ginx.WithHandlerName("DownloadPodLogs"))

	reg.Register(r, &DownloadPodFileTransfer{}, ginx.WithHandlerName("DownloadPodFile"))

	reg.Register(r, &DownloadPodFileV2Transfer{}, ginx.WithHandlerName("DownloadPodFileV2"))

	reg.Register(r, &PortforwardTransfer{}, ginx.WithHandlerName("Portforward"))

}
