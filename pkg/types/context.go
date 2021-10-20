package types

type Context struct {
	Config *MyConf

	CoreLog CoreLog

	Log Logger

	Tracer Tracer

	GinX GinXer

	Storage Storage

	Cache Cache

	Apis Apis

	Version Version

	localVersion int

	Reuse func(x interface{})
}
