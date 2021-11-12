package core

import (
	"fmt"
	"github.com/fitan/magic/pkg/ginx"
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
)

func GinXHandlerRegister(i gin.IRouter, transfer types.GinXTransfer, o ...ginx.GinXHandlerOption) {
	i.Handle(transfer.Method(), transfer.Url(), func(c *gin.Context) {
		core := GetCorePool().GetObj()
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
				core.Log = core.CoreLog.TraceLog(core.GinX.GinCtx().GetString(ginx.FnName))
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
	})
}

type ginXRegister struct {
	EntryMid []types.Middleware
}

func (g *ginXRegister) Reload(c *types.Core) {
}

func (g *ginXRegister) With(o ...types.Option) types.Register {
	return g
}

func (g *ginXRegister) Set(c *types.Core) {
	c.GinX = ginx.NewGin(ginx.WithEntryMid(&g.EntryMid))
}

func (g *ginXRegister) Unset(c *types.Core) {
	c.GinX.SetBindReq(nil)
	c.GinX.SetBindRes(nil)
	c.GinX.SetBindErr(nil)
	c.GinX.SetHandlerMid(nil)
	c.GinX.SetGinCtx(nil)
}
