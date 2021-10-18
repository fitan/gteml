package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)

	userRoute(r)

	return r
}
