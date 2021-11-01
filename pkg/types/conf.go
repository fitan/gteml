package types

type MyConf struct {
	Log   Log   `yaml:"log"`
	Trace Trace `yaml:"trace"`
	Api   Api   `yaml:"api"`
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	App   App   `yaml:"app"`
}

type Api struct {
	Baidu  Baidu  `yaml:"baidu"`
	Taobao Taobao `yaml:"taobao"`
}

type Baidu struct {
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
}

type Taobao struct {
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
}

type Mysql struct {
	Url string `yaml:"url"`
}

type Trace struct {
	Open               bool   `yaml:"open"`
	TracerProviderAddr string `yaml:"tracerProviderAddr"`
}

type Log struct {
	Lervel      int    `yaml:"lervel"`
	TraceLervel int    `yaml:"traceLervel"`
	FileName    string `yaml:"fileName"`
}

type Redis struct {
	Url       string `yaml:"url"`
	Password  string `yaml:"password"`
	Db        int    `yaml:"db"`
	OpenTrace bool   `yaml:"openTrace"`
}

type App struct {
	Name string `yaml:"name"`
}
