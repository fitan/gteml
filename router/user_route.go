package router

import (
	"github.com/fitan/magic/gen/transfer/user"
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func userRoute(r gin.IRouter) {
	core.GinXHandlerRegister(r, &user.CreateTransfer{}, ginx.WithHandlerName("get user"), ginx.WithHandlerMid(&ginx.TestMid{}))
	core.GinXHandlerRegister(r, &user.SayHelloTransfer{}, ginx.WithHandlerName("say hello"))
}
