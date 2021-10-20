package types

import "github.com/gin-gonic/gin"

type GinXer interface {
	BindTransfer(core *Context, i GinXBinder)
	SetGinCtx(c *gin.Context)
	GinCtx() *gin.Context
	SetBindReq(interface{})
	SetBindRes(interface{})
	SetBindErr(error)
}

type GinXBinder interface {
	BindVal(c *Context) (interface{}, error)
	BindFn(c *Context) (interface{}, error)
}

type GinXTransfer interface {
	Method() string
	Url() string
	Binder() GinXBinder
}
