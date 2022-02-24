package types

type MyConf struct {
	Redis     Redis     `yaml:"redis"`
	Jwt       Jwt       `yaml:"jwt"`
	K8sConf   K8sConf   `yaml:"k8sConf"`
	App       App       `yaml:"app"`
	Pyroscope Pyroscope `yaml:"pyroscope"`
	Apis      Apis      `yaml:"apis"`
	Mysql     Mysql     `yaml:"mysql"`
	Rbac      Rbac      `yaml:"rbac"`
	Swagger   Swagger   `yaml:"swagger"`
	Consul    Consul    `yaml:"consul"`
	Log       Log       `yaml:"Logger"`
	Trace     Trace     `yaml:"trace"`
}

type Redis struct {
	Url       string `yaml:"url"`
	Password  string `yaml:"password"`
	Db        int    `yaml:"db"`
	OpenTrace bool   `yaml:"openTrace"`
}

type Pyroscope struct {
	Open bool   `yaml:"open"`
	Url  string `yaml:"url"`
}

type Jwt struct {
	TokenHeadName string `yaml:"tokenHeadName"`
	Realm         string `yaml:"realm"`
	IdentityKey   string `yaml:"identityKey"`
	SecretKey     string `yaml:"secretKey"`
	Timeout       string `yaml:"timeout"`
	MaxRefresh    string `yaml:"maxRefresh"`
}

type Baidu struct {
	RestyDebug bool   `yaml:"restyDebug"`
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
}

type Swagger struct {
	Enable bool `yaml:"enable"`
}

type Trace struct {
	Open               bool   `yaml:"open"`
	TracerProviderAddr string `yaml:"tracerProviderAddr"`
}

type K8sConf struct {
	ConfigPath string `yaml:"configPath"`
}

type Apis struct {
	Baidu  Baidu  `yaml:"baidu"`
	Taobao Taobao `yaml:"taobao"`
}

type Taobao struct {
	Url        string `yaml:"url"`
	TraceDebug bool   `yaml:"traceDebug"`
	RestyDebug bool   `yaml:"restyDebug"`
}

type Mysql struct {
	ConnMaxIdleTime string `yaml:"ConnMaxIdleTime"`
	Url             string `yaml:"url"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime string `yaml:"connMaxLifetime"`
}

type Rbac struct {
	Model string `yaml:"model"`
}

type Consul struct {
	Addr string `yaml:"addr"`
}

type App struct {
	Addr string `yaml:"addr"`
	Name string `yaml:"name"`
}

type Log struct {
	Lervel      int    `yaml:"lervel"`
	Dir         string `yaml:"dir"`
	TraceLervel int    `yaml:"traceLervel"`
	FileName    string `yaml:"fileName"`
}
