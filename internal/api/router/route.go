package router

import (
	"github.com/fitan/gteml/internal/api/gen/transfer/user"
	"github.com/fitan/gteml/pkg/core"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)

	core.GinXHandlerRegister(r, &user.CreateTransfer{}, core.WithHandlerName("get user"))

	return r
}
