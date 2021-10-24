package core

import (
	"github.com/fitan/gteml/pkg/api"
	"github.com/fitan/gteml/pkg/types"
)

//type ContextPool struct {
//	p            sync.Pool
//	registerList []Register
//}
//
//func ConfReloadHook() {
//}
//
//var registerList []Register
//
//func InsetRegister(os ...Register) {
//	registerList = append(registerList, os...)
//}

func init() {
	GetCtxPool().RegisterList([]types.Register{
		&ConfReg{},
		&logRegister{},
		&ginXRegister{},
		&Trace{},
		&api.ApisRegister{},
		&VersionReg{},
		&PoolReg{},
	})
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
