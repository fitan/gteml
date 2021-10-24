package types

type MyConf struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	App   App   `yaml:"app"`
	Log   Log   `yaml:"log"`
	Trace Trace `yaml:"trace"`
	Api   Api   `yaml:"api"`
}

type Mysql struct {
	Url string `yaml:"url"`
}

type Redis struct {
	Url string `yaml:"url"`
}

type App struct {
	Name string `yaml:"name"`
}

type Log struct {
	Lervel   int    `yaml:"lervel"`
	FileName string `yaml:"fileName"`
}

type Api struct {
	Baidu  Baidu  `yaml:"baidu"`
	Taobao Taobao `yaml:"taobao"`
}

type Taobao struct {
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
	Url        string `yaml:"url"`
}

type Trace struct {
	Open bool `yaml:"open"`
}

type Baidu struct {
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
}
