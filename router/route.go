package router

import (
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/prometheus"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(core.SetCore())
	prometheus.UseGinprom(r)
	pprof.Register(r)

	registerRouter(r)

	return r
}

// 注册路由
func registerRouter(r *gin.Engine) {
	gReg := core.NewGinXHandlerRegister()
	userRoute(r, gReg)
}
