package ginx

import (
	"crypto/md5"
	"fmt"
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"path"
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

func (g *GinXHandlerRegister) Group(options ...GinXHandlerOption) *GinXHandlerRegister {
	os := make([]GinXHandlerOption, 0, len(g.options)+len(options))
	os = append(os, g.options...)
	os = append(os, options...)
	return NewGinXHandlerRegister(os...)
}

var CollectRouterSlice [][]string

func CollectRouter(i interface{}, transfer types.GinXTransfer) {
	if g, ok := i.(*gin.RouterGroup); ok {
		md5 := md5.Sum([]byte(path.Join(transfer.FuncName(), transfer.Method(), g.BasePath(), transfer.Url())))
		CollectRouterSlice = append(CollectRouterSlice, []string{transfer.FuncName(), transfer.Method(), path.Join(g.BasePath(), transfer.Url()), cast.ToString(md5)})
		log.Printf("method: %v, path: %v, md5: %x", transfer.Method(), path.Join(g.BasePath(), transfer.Url()), md5)
	} else {
		md5 := md5.Sum([]byte(path.Join(transfer.FuncName(), transfer.Method(), transfer.Url())))
		CollectRouterSlice = append(CollectRouterSlice, []string{transfer.FuncName(), transfer.Method(), transfer.Url(), cast.ToString(md5)})
		log.Printf("method: %v, path: %v, md5: %x", transfer.Method(), transfer.Url(), md5)
	}
}

func ginXHandlerRegister(i gin.IRouter, transfer types.GinXTransfer, o ...GinXHandlerOption) {
	CollectRouter(i, transfer)
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
			// 加载中间件option
			for _, f := range o {
				err := f(core)
				if err != nil {
					core.GinX.SetError(fmt.Errorf("load option err: %w", err))
					break
				}
			}

			if core.Tracer.IsOpen() {

				if core.CoreLog.IsOpenTrace() {
					// 设置tracelog
					core.Log = core.CoreLog.TraceLog(core.GinX.GinCtx().GetString(FnName))
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
