package types

type Core struct {
	Config *MyConf

	CoreLog CoreLoger

	Log Logger

	Tracer Tracer

	GinX GinXer

	Storage Storage

	Cache Cache

	Apis Apis

	Prom Promer

	Version Version

	LocalVersion int

	Pool Pooler
}
