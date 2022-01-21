package types

type Core struct {
	Config Confer

	CoreLog CoreLoger

	Log Logger

	Tracer Tracer

	GinX GinXer

	//Storage Storage

	Dao DAOer

	Services Serviceser

	Apis Apis

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

func (c *Core) GetServices() Serviceser {
	return c.Services
}

func (c *Core) GetApis() Apis {
	return c.Apis
}

func (c *Core) GetProm() Promer {
	return c.Prom
}

func (c *Core) GetDao() DAOer {
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
	GetServices() Serviceser
	GetApis() Apis
	GetProm() Promer
	GetDao() DAOer
}
