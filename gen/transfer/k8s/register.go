package k8s

import (
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter, reg *ginx.GinXHandlerRegister) {

	reg.Register(r, &GetAppTransfer{}, ginx.WithHandlerName("GetApp"))

	reg.Register(r, &CreateWorkerTransfer{}, ginx.WithHandlerName("CreateWorker"))

	reg.Register(r, &GetPodsTransfer{}, ginx.WithHandlerName("GetPods"))

}
