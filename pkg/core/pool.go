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
	InsetRegister(&Confer{})
	InsetRegister(&logRegister{})
	InsetRegister(&ginXRegister{})
	InsetRegister(&Trace{})
	InsetRegister(&ApisRegister{})
	InsetRegister(&VersionReg{})
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
	// 初始化对象版本号
	core.SetLocalVersion()
	return core
	//return NewCore().Init()
}

var corePool = sync.Pool{
	New: NewPoolFn,
}

func GetCore() *Context {
	return corePool.Get().(*Context)
}

func PutCore(x interface{}) {
	corePool.Put(x)
}
