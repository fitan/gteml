package core

import (
	"github.com/fitan/gteml/pkg/conf"
	"log"
)

// Redis
type Redis struct {
	Url string `yaml:"url"`
}

// Yaml2Go
type MyConf struct {
	Trace OpenTrace `yaml:"trace"`
	Api   Api       `yaml:"api"`
	Mysql Mysql     `yaml:"mysql"`
	Redis Redis     `yaml:"redis"`
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

var myConf *MyConf

func init() {
	myConf = &MyConf{}
	w, err := conf.WatchFile("conf", []string{"./"}, myConf)
	if err != nil {
		panic(err)
	}
	c := w.GetSignal()
	go func() {
		for {
			<-c
			log.Println("reload config")
		}
	}()
}

func (c *Confer) With(o ...Option) Register {
	panic("implement me")
}

func (c *Confer) Set(ctx *Context) {
	ctx.Config = myConf
}

func (c *Confer) Unset(ctx *Context) {
}
