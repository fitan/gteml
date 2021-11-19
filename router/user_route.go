package router

import (
	"github.com/fitan/magic/gen/transfer/user"
	"github.com/fitan/magic/pkg/core"
	"github.com/gin-gonic/gin"
)

func userRoute(r gin.IRouter, reg *core.GinXHandlerRegister) {
	user.Register(r, reg)
}
