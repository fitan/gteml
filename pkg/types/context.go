package types

type Context struct {
	Config *MyConf

	CoreLog CoreLoger

	Log Logger

	Tracer Tracer

	GinX GinXer

	Storage Storage

	Cache Cache

	Apis Apis

	Version Version

	LocalVersion int

	Pool Pooler
}
