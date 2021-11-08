package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	pprof.Register(r)

	registerRouter(r)

	return r
}

// 注册路由
func registerRouter(r *gin.Engine) {
	userRoute(r)
}
