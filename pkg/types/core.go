package types

import (
	apiTypes "github.com/fitan/magic/apis/types"
	daoTypes "github.com/fitan/magic/dao/types"
	serviceTypes "github.com/fitan/magic/services/types"
)

type Core struct {
	Config Confer

	CoreLog CoreLoger

	Log Logger

	Tracer Tracer

	GinX GinXer

	//Storage Storage

	Dao daoTypes.DAOer

	Services serviceTypes.Serviceser

	Apis apiTypes.Apis

	Prom Promer

	Version Version

	LocalVersion int

	Pool Pooler
}

func (c *Core) GetConfig() Confer {
	return c.Config
}

func (c *Core) GetCoreLog() CoreLoger {
	return c.CoreLog
}

func (c *Core) GetServices() serviceTypes.Serviceser {
	return c.Services
}

func (c *Core) GetApis() apiTypes.Apis {
	return c.Apis
}

func (c *Core) GetProm() Promer {
	return c.Prom
}

func (c *Core) GetDao() daoTypes.DAOer {
	return c.Dao
}

func (c *Core) GetTrace() Tracer {
	return c.Tracer
}

func (c *Core) GetGinX() GinXer {
	return c.GinX
}

type TracerCore interface {
	GetTrace() Tracer
}

type ServiceCore interface {
	GetTrace() Tracer
	GetConfig() Confer
	GetCoreLog() CoreLoger
	GetServices() serviceTypes.Serviceser
	GetApis() apiTypes.Apis
	GetProm() Promer
	GetDao() daoTypes.DAOer
}
