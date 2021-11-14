package types

type Core struct {
	Config Confer

	CoreLog CoreLoger

	Log Logger

	Tracer Tracer

	GinX GinXer

	Storage Storage

	Cache Cache

	Services ServicesI

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

func (c *Core) GetStorage() Storage {
	return c.Storage
}

func (c *Core) GetCache() Cache {
	return c.Cache
}

func (c *Core) GetServices() ServicesI {
	return c.Services
}

func (c *Core) GetApis() Apis {
	return c.Apis
}

func (c *Core) GetProm() Promer {
	return c.Prom
}

type ServiceCore interface {
	GetConfig() Confer
	GetCoreLog() CoreLoger
	GetStorage() Storage
	GetCache() Cache
	GetServices() ServicesI
	GetApis() Apis
	GetProm() Promer
}
