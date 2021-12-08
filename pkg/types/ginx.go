package types

import "github.com/gin-gonic/gin"

type GinXer interface {
	BindTransfer(core *Core, i GinXBinder)
	SetGinCtx(ctx *gin.Context)
	GinCtx() *gin.Context
	SetRequest(interface{})
	Request() interface{}
	SetResponse(interface{})
	Response() interface{}
	SetError(error)
	Errors() []error
	LastError() error
	Reset()
	SetEntryMid(m *[]Middleware)
	SetHandlerMid(m *[]Middleware)
}

type GinXBinder interface {
	BindVal(core *Core) (interface{}, error)
	BindFn(core *Core) (interface{}, error)
}

type GinXTransfer interface {
	Method() string
	Url() string
	Binder() GinXBinder
}

type Option func(core *Core)

type Middleware interface {
	BindValBefor(core *Core) bool
	BindValAfter(core *Core) bool
	BindFnAfter(core *Core) bool
	Forever(core *Core)
}
