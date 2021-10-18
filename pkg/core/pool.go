package core

import (
	"sync"
)

type ContextPool struct {
	p            sync.Pool
	registerList []Register
}

func ConfReloadHook() {
}

var registerList []Register

func InsetRegister(os ...Register) {
	registerList = append(registerList, os...)
}

func init() {
	InsetRegister(&logRegister{})
	InsetRegister((&ginXRegister{}).With())
	InsetRegister(&Trace{})
	InsetRegister(&ApisRegister{})
}

//func NewCore() *Context {
//	core := &Context{}
//
//	core.Register(&logRegister{})
//	core.Register((&ginXRegister{}).With())
//	core.Register(&TraceRegister{})
//	core.Register(&ApisRegister{})
//	return core
//}

func NewPoolFn() interface{} {
	core := &Context{}
	for _, r := range registerList {
		r.Set(core)
	}
	return core
	//return NewCore().Init()
}

var corePool = sync.Pool{
	New: NewPoolFn,
}

func GetCore() *Context {
	return corePool.Get().(*Context).Init()
}

func PutCore(x interface{}) {
	corePool.Put(x)
}
