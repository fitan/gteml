package user

import (
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter) {

	core.GinXHandlerRegister(r, &CreateTransfer{}, ginx.WithHandlerName("Create"))

	core.GinXHandlerRegister(r, &SayHelloTransfer{}, ginx.WithHandlerName("SayHello"))

}
