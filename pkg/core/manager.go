package core

import (
	"github.com/fitan/magic/pkg/api"
	"github.com/fitan/magic/pkg/types"
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

var Conf *ConfReg

func init() {
	Conf = NewConfReg()
	GetCtxPool().RegisterList([]types.Register{
		Conf,
		&Trace{},
		&logRegister{},
		&storageReg{},
		&ginXRegister{},
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
