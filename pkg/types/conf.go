package types

// Redis
type Redis struct {
	Url string `yaml:"url"`
}

// Yaml2Go
type MyConf struct {
	App   App       `yaml:"app"`
	Log   Log       `yaml:"log"`
	Trace OpenTrace `yaml:"trace"`
	Api   Api       `yaml:"api"`
	Mysql Mysql     `yaml:"mysql"`
	Redis Redis     `yaml:"redis"`
}

type Log struct {
	FileName string `yaml:"fileName"`
	Lervel   int    `yaml:"lervel"`
}

type App struct {
	Name string `yaml:"name"`
}

// Trace
type OpenTrace struct {
	Open bool `yaml:"open"`
}

// Api
type Api struct {
	Baidu  Baidu  `yaml:"baidu"`
	Taobao Taobao `yaml:"taobao"`
}

// Baidu
type Baidu struct {
	Url string `yaml:"url"`
}

// Taobao
type Taobao struct {
	Url string `yaml:"url"`
}

// Mysql
type Mysql struct {
	Url string `yaml:"url"`
}

type Confer struct {
	MyConf
}
