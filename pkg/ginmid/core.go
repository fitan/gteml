package ginmid

import (
	core "github.com/fitan/magic/pkg/core"
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
)

func SetCore() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := core.GetCorePool().GetObj()
		c.GetGinX().SetGinCtx(ctx)
		c.Tracer.SetCtx(ctx.Request.Context())
		ctx.Set(types.CoreKey, c)
		ctx.Next()
		c.Pool.ReUse(c)
	}
}
