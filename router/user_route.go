package router

import (
	"github.com/fitan/magic/gen/transfer/user"
	"github.com/gin-gonic/gin"
)

func userRoute(r gin.IRouter) {
	user.Register(r)
	//core.GinXHandlerRegister(r, &user_back.CreateTransfer{}, ginx.WithHandlerName("get user"), ginx.WithHandlerMid(&ginx.TestMid{}))
	//core.GinXHandlerRegister(r, &user_back.SayHelloTransfer{}, ginx.WithHandlerName("say hello"))
}
