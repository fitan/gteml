package ginmid

import (
	"github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
)

func SetCore() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		core := core.GetCorePool().GetObj()
		ctx.Set(types.CoreKey, core)
		ctx.Next()
		core.Pool.ReUse(core)
	}
}
