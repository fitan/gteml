package router

import (
	_ "github.com/fitan/magic/docs"
	core2 "github.com/fitan/magic/pkg/core"
	"github.com/gin-gonic/gin"
)
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"

func swag(i gin.IRouter) {
	core := core2.GetCorePool().GetObj()
	if core.GetConfig().Swagger.Enable {
		i.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
