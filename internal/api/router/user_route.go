package router

import (
	"github.com/fitan/gteml/internal/api/gen/transfer/user"
	"github.com/fitan/gteml/pkg/core"
	"github.com/gin-gonic/gin"
)

func userRoute(r gin.IRouter) {
	core.GinXHandlerRegister(r, &user.CreateTransfer{}, core.WithHandlerName("get user"))
}
