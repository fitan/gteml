package types

type MyConf struct {
	Redis     Redis     `yaml:"redis"`
	Jwt       Jwt       `yaml:"jwt"`
	Swagger   Swagger   `yaml:"swagger"`
	App       App       `yaml:"app"`
	Log       Log       `yaml:"log"`
	Api       Api       `yaml:"api"`
	Mysql     Mysql     `yaml:"mysql"`
	Trace     Trace     `yaml:"trace"`
	Pyroscope Pyroscope `yaml:"pyroscope"`
	Rbac      Rbac      `yaml:"rbac"`
}

type Jwt struct {
	SecretKey     string `yaml:"secretKey"`
	Timeout       string `yaml:"timeout"`
	MaxRefresh    string `yaml:"maxRefresh"`
	TokenHeadName string `yaml:"tokenHeadName"`
	Realm         string `yaml:"realm"`
	IdentityKey   string `yaml:"identityKey"`
}

type Swagger struct {
	Enable bool `yaml:"enable"`
}

type App struct {
	Name string `yaml:"name"`
}

type Log struct {
	Lervel      int    `yaml:"lervel"`
	Dir         string `yaml:"dir"`
	TraceLervel int    `yaml:"traceLervel"`
	FileName    string `yaml:"fileName"`
}

type Baidu struct {
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

type Trace struct {
	Open               bool   `yaml:"open"`
	TracerProviderAddr string `yaml:"tracerProviderAddr"`
}

type Redis struct {
	Password  string `yaml:"password"`
	Db        int    `yaml:"db"`
	OpenTrace bool   `yaml:"openTrace"`
	Url       string `yaml:"url"`
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
	RestyDebug bool   `yaml:"restyDebug"`
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
}

type Rbac struct {
	Model string `yaml:"model"`
}
