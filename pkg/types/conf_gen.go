package types

type MyConf struct {
	K8sConf   K8sConf   `yaml:"k8sConf"`
	App       App       `yaml:"app"`
	Log       Log       `yaml:"log"`
	Pyroscope Pyroscope `yaml:"pyroscope"`
	Api       Api       `yaml:"api"`
	Mysql     Mysql     `yaml:"mysql"`
	Redis     Redis     `yaml:"redis"`
	Trace     Trace     `yaml:"trace"`
	Jwt       Jwt       `yaml:"jwt"`
	Rbac      Rbac      `yaml:"rbac"`
	Swagger   Swagger   `yaml:"swagger"`
}

type Log struct {
	Lervel      int    `yaml:"lervel"`
	Dir         string `yaml:"dir"`
	TraceLervel int    `yaml:"traceLervel"`
	FileName    string `yaml:"fileName"`
}

type Redis struct {
	Password  string `yaml:"password"`
	Db        int    `yaml:"db"`
	OpenTrace bool   `yaml:"openTrace"`
	Url       string `yaml:"url"`
}

type Trace struct {
	Open               bool   `yaml:"open"`
	TracerProviderAddr string `yaml:"tracerProviderAddr"`
}

type Jwt struct {
	Realm         string `yaml:"realm"`
	IdentityKey   string `yaml:"identityKey"`
	SecretKey     string `yaml:"secretKey"`
	Timeout       string `yaml:"timeout"`
	MaxRefresh    string `yaml:"maxRefresh"`
	TokenHeadName string `yaml:"tokenHeadName"`
}

type Swagger struct {
	Enable bool `yaml:"enable"`
}

type App struct {
	Name string `yaml:"name"`
}

type Taobao struct {
	RestyDebug bool   `yaml:"restyDebug"`
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
}

type Mysql struct {
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime string `yaml:"connMaxLifetime"`
	ConnMaxIdleTime string `yaml:"ConnMaxIdleTime"`
	Url             string `yaml:"url"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
}

type Rbac struct {
	Model string `yaml:"model"`
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

type Pyroscope struct {
	Open bool   `yaml:"open"`
	Url  string `yaml:"url"`
}

type K8sConf struct {
	ConfigPath string `yaml:"configPath"`
}
