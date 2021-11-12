package router

import (
	"github.com/fitan/magic/pkg/prometheus"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	prometheus.UseGinprom(r)

	pprof.Register(r)

	registerRouter(r)
	r.GET("/hello", func(context *gin.Context) {

	})

	return r
}

// 注册路由
func registerRouter(r *gin.Engine) {
	userRoute(r)
}
