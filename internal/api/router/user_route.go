package router

import (
	"github.com/fitan/magic/internal/api/gen/transfer/user"
	"github.com/fitan/magic/pkg/context"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/gin-gonic/gin"
)

func userRoute(r gin.IRouter) {
	context.GinXHandlerRegister(r, &user.CreateTransfer{}, ginx.WithHandlerName("get user"))
}
