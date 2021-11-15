package user

import (
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func Register(r gin.IRouter, reg *core.GinXHandlerRegister) {

	reg.Register(r, &CreateTransfer{}, ginx.WithHandlerName("Create"))
	reg.Register(r, &SayHelloTransfer{}, ginx.WithHandlerName("SayHello"))

}
