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
	GetCorePool().RegisterList([]types.Register{
		Conf,
		&Trace{},
		&logRegister{},
		&storageReg{},
		&CacheReg{},
		&ginXRegister{},
		&api.ApisRegister{},
		&VersionReg{},
		&PoolReg{},
	})
}

//func NewCore() *Core {
//	core := &Core{}
//
//	core.Register(&logRegister{})
//	core.Register((&ginXRegister{}).With())
//	core.Register(&TraceRegister{})
//	core.Register(&ApisRegister{})
//	return core
//}
