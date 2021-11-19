package types

type MyConf struct {
	Jwt       Jwt       `yaml:"jwt"`
	App       App       `yaml:"app"`
	Log       Log       `yaml:"log"`
	Trace     Trace     `yaml:"trace"`
	Pyroscope Pyroscope `yaml:"pyroscope"`
	Api       Api       `yaml:"api"`
	Mysql     Mysql     `yaml:"mysql"`
	Redis     Redis     `yaml:"redis"`
}

type App struct {
	Name string `yaml:"name"`
}

type Log struct {
	TraceLervel int    `yaml:"traceLervel"`
	FileName    string `yaml:"fileName"`
	Lervel      int    `yaml:"lervel"`
	Dir         string `yaml:"dir"`
}

type Baidu struct {
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
	Url        string `yaml:"url"`
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

type Mysql struct {
	Url             string `yaml:"url"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime string `yaml:"connMaxLifetime"`
	ConnMaxIdleTime string `yaml:"ConnMaxIdleTime"`
}

type Redis struct {
	OpenTrace bool   `yaml:"openTrace"`
	Url       string `yaml:"url"`
	Password  string `yaml:"password"`
	Db        int    `yaml:"db"`
}

type Jwt struct {
	MaxRefresh    string `yaml:"maxRefresh"`
	TokenHeadName string `yaml:"tokenHeadName"`
	Realm         string `yaml:"realm"`
	IdentityKey   string `yaml:"identityKey"`
	SecretKey     string `yaml:"secretKey"`
	Timeout       string `yaml:"timeout"`
}

type Trace struct {
	Open               bool   `yaml:"open"`
	TracerProviderAddr string `yaml:"tracerProviderAddr"`
}
