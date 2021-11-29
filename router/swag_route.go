package router

import (
	_ "github.com/fitan/magic/docs"
	"github.com/fitan/magic/pkg/core"
	"github.com/gin-gonic/gin"
)
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"

func SwagRoute(i gin.IRouter) {
	core := core.GetCorePool().GetObj()
	if core.GetConfig().GetMyConf().Swagger.Enable {
		i.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
