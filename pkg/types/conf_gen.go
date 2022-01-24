package types

type MyConf struct {
	App       App       `yaml:"app"`
	Apis      Apis      `yaml:"apis"`
	Mysql     Mysql     `yaml:"mysql"`
	Swagger   Swagger   `yaml:"swagger"`
	Rbac      Rbac      `yaml:"rbac"`
	Consul    Consul    `yaml:"consul"`
	Log       Log       `yaml:"log"`
	Trace     Trace     `yaml:"trace"`
	Pyroscope Pyroscope `yaml:"pyroscope"`
	Redis     Redis     `yaml:"redis"`
	Jwt       Jwt       `yaml:"jwt"`
	K8sConf   K8sConf   `yaml:"k8sConf"`
}

type Swagger struct {
	Enable bool `yaml:"enable"`
}

type Pyroscope struct {
	Open bool   `yaml:"open"`
	Url  string `yaml:"url"`
}

type K8sConf struct {
	ConfigPath string `yaml:"configPath"`
}

type Trace struct {
	Open               bool   `yaml:"open"`
	TracerProviderAddr string `yaml:"tracerProviderAddr"`
}

type App struct {
	Name string `yaml:"name"`
}

type Baidu struct {
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
}

type Consul struct {
	Addr string `yaml:"addr"`
}

type Log struct {
	Lervel      int    `yaml:"lervel"`
	Dir         string `yaml:"dir"`
	TraceLervel int    `yaml:"traceLervel"`
	FileName    string `yaml:"fileName"`
}

type Mysql struct {
	Url             string `yaml:"url"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime string `yaml:"connMaxLifetime"`
	ConnMaxIdleTime string `yaml:"ConnMaxIdleTime"`
}

type Rbac struct {
	Model string `yaml:"model"`
}

type Redis struct {
	OpenTrace bool   `yaml:"openTrace"`
	Url       string `yaml:"url"`
	Password  string `yaml:"password"`
	Db        int    `yaml:"db"`
}

type Jwt struct {
	TokenHeadName string `yaml:"tokenHeadName"`
	Realm         string `yaml:"realm"`
	IdentityKey   string `yaml:"identityKey"`
	SecretKey     string `yaml:"secretKey"`
	Timeout       string `yaml:"timeout"`
	MaxRefresh    string `yaml:"maxRefresh"`
}

type Apis struct {
	Taobao Taobao `yaml:"taobao"`
	Baidu  Baidu  `yaml:"baidu"`
}

type Taobao struct {
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
}
