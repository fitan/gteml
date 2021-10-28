package types

type MyConf struct {
	Log   Log   `yaml:"log"`
	Trace Trace `yaml:"trace"`
	Api   Api   `yaml:"api"`
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	App   App   `yaml:"app"`
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

type Redis struct {
	Db        int    `yaml:"db"`
	OpenTrace bool   `yaml:"openTrace"`
	Url       string `yaml:"url"`
	Password  string `yaml:"password"`
}

type App struct {
	Name string `yaml:"name"`
}

type Api struct {
	Baidu  Baidu  `yaml:"baidu"`
	Taobao Taobao `yaml:"taobao"`
}

type Log struct {
	Lervel   int    `yaml:"lervel"`
	FileName string `yaml:"fileName"`
}

type Trace struct {
	Open               bool   `yaml:"open"`
	TracerProviderAddr string `yaml:"tracerProviderAddr"`
}

type Mysql struct {
	Url string `yaml:"url"`
}
