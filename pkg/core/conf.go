package core

import (
	"github.com/fitan/gteml/pkg/conf"
	"github.com/fitan/gteml/pkg/types"
	"log"
	"runtime"
	"time"
)

var myConf *types.MyConf

func init() {
	myConf = &types.MyConf{}
	w, err := conf.WatchFile("conf", []string{"./"}, myConf, 5*time.Second)
	if err != nil {
		panic(err)
	}
	c := w.GetSignal()
	go func() {
		for {
			<-c
			GetCtxPool().GetObj().Version.AddVersion()
			//配置文件reload后 gc触发清理pool中的对象
			runtime.GC()
			log.Println("reload config version: ", GetCtxPool().GetObj().Version.Version())
		}
	}()
}

type ConfReg struct {
}

func (c *ConfReg) With(o ...types.Option) types.Register {
	return c
}

func (c *ConfReg) Reload(ctx *types.Context) {
}

func (c *ConfReg) Set(ctx *types.Context) {
	ctx.Config = myConf
}

func (c *ConfReg) Unset(ctx *types.Context) {
}

//func (c *ConfReg) With(o ...types.Option) Register {
//	panic("implement me")
//}
//
//func (c *ConfReg) Reload(ctx *Context) {
//}
//
//func (c *ConfReg) Set(ctx *Context) {
//	ctx.Config = myConf
//}
//
//func (c *ConfReg) Unset(ctx *Context) {
//}
