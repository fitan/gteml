package ginx

import (
	"fmt"
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinXHandlerRegister struct {
	options []GinXHandlerOption
}

func NewGinXHandlerRegister(options ...GinXHandlerOption) *GinXHandlerRegister {
	return &GinXHandlerRegister{options: options}
}

func (g *GinXHandlerRegister) Register(i gin.IRouter, transfer types.GinXTransfer, o ...GinXHandlerOption) {
	option := make([]GinXHandlerOption, 0, len(g.options)+len(o))
	option = append(option, g.options...)
	option = append(option, o...)
	ginXHandlerRegister(i, transfer, option...)

}

func ginXHandlerRegister(i gin.IRouter, transfer types.GinXTransfer, o ...GinXHandlerOption) {
	i.Handle(
		transfer.Method(), transfer.Url(), func(c *gin.Context) {
			var core *types.Core
			if value, ok := c.Get(types.CoreKey); ok {
				core = value.(*types.Core)
			} else {
				c.JSON(http.StatusInternalServerError, GinXResult{
					Code: 5000,
					Msg:  "gin ctx not found types.Core",
					Data: nil,
				})
				return
				//core = coreSrc.GetCorePool().GetObj()
			}
			//gin的request ctx放到trace里
			//core.SetCtx(c.Request.Core())
			// core包裹gin context
			core.GinX.SetGinCtx(c)
			core.Tracer.SetCtx(c.Request.Context())
			// 加载中间件option
			for _, f := range o {
				err := f(core)
				if err != nil {
					core.GinX.SetBindErr(fmt.Errorf("load option err: %w", err))
					break
				}
			}

			if core.Tracer.IsOpen() {

				if core.CoreLog.IsOpenTrace() {
					// 设置tracelog
					core.Log = core.CoreLog.ApmLog(core.GinX.GinCtx().GetString(FnName))
					// 如果打开trace则end
					defer core.Log.Sync()
				} else {
					core.Log = core.CoreLog.Log()
				}
			} else {
				// 普通log
				core.Log = core.CoreLog.Log()
			}

			core.GinX.BindTransfer(core, transfer.Binder())
		},
	)
}
