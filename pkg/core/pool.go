package core

import (
	"sync"
)

func NewCore() *Context {
	core := &Context{}

	core.Register(&logRegister{})
	core.Register(&ginXRegister{})
	core.releaseFn = PutCore
	return core
}

var corePool = sync.Pool{
	New: NewPoolFn,
}

func GetCore() interface{} {
	return corePool.Get()
}

func PutCore(x interface{}) {
	corePool.Put(x)
}

func NewPoolFn() interface{} {
	c := NewCore()
	c.Init()
	return c
}
