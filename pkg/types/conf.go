package types

type MyConf struct {
	Api       Api       `yaml:"api"`
	Mysql     Mysql     `yaml:"mysql"`
	Redis     Redis     `yaml:"redis"`
	App       App       `yaml:"app"`
	Log       Log       `yaml:"log"`
	Trace     Trace     `yaml:"trace"`
	Pyroscope Pyroscope `yaml:"pyroscope"`
}

type Baidu struct {
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
}

type Mysql struct {
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime string `yaml:"connMaxLifetime"`
	ConnMaxIdleTime string `yaml:"ConnMaxIdleTime"`
	Url             string `yaml:"url"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
}

type Redis struct {
	Url       string `yaml:"url"`
	Password  string `yaml:"password"`
	Db        int    `yaml:"db"`
	OpenTrace bool   `yaml:"openTrace"`
}

type Log struct {
	Lervel      int    `yaml:"lervel"`
	Dir         string `yaml:"dir"`
	TraceLervel int    `yaml:"traceLervel"`
	FileName    string `yaml:"fileName"`
}

type Pyroscope struct {
	Open bool   `yaml:"open"`
	Url  string `yaml:"url"`
}

type Api struct {
	Baidu  Baidu  `yaml:"baidu"`
	Taobao Taobao `yaml:"taobao"`
}

type Taobao struct {
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
}

type App struct {
	Name string `yaml:"name"`
}

type Trace struct {
	Open               bool   `yaml:"open"`
	TracerProviderAddr string `yaml:"tracerProviderAddr"`
}
