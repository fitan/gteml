package router

import (
	"github.com/fitan/gteml/pkg/core"
	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.New()
	r.GET("/user", core.GinXHandlerRegister())
}
