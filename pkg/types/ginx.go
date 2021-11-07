package types

import "github.com/gin-gonic/gin"

type GinXer interface {
	BindTransfer(core *Core, i GinXBinder)
	SetGinCtx(c *gin.Context)
	GinCtx() *gin.Context
	SetBindReq(interface{})
	BindReq() interface{}
	SetBindRes(interface{})
	BindRes() interface{}
	SetBindErr(error)
	BindErr() error
	Result(c *Core)
}

type GinXBinder interface {
	BindVal(c *Core) (interface{}, error)
	BindFn(c *Core) (interface{}, error)
}

type GinXTransfer interface {
	Method() string
	Url() string
	Binder() GinXBinder
}

type Option func(c *Core)
