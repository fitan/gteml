package types

import "github.com/gin-gonic/gin"

type GinXer interface {
	BindTransfer(core *Context, i GinXBinder)
	SetGinCtx(c *gin.Context)
	GinCtx() *gin.Context
	SetBindReq(interface{})
	BindReq() interface{}
	SetBindRes(interface{})
	BindRes() interface{}
	SetBindErr(error)
	BindErr() error
	Result(c *Context)
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

type Option func(c *Context)
