package core

import (
	"github.com/fitan/magic/pkg/conf"
	"github.com/fitan/magic/pkg/types"
	"log"
	"runtime"
	"time"
)

type Conf struct {
	myConf *types.MyConf
}

func (c *Conf) GetMyConf() *types.MyConf {
	return c.myConf
}

func NewConfReg() *ConfRegister {
	myConf := &types.MyConf{}
	w, err := conf.WatchFile("conf", []string{"./"}, myConf, 5*time.Second)
	if err != nil {
		panic(err)
	}
	c := w.GetSignal()
	go func() {
		for {
			<-c
			GetCorePool().Reload()
			GetCorePool().GetObj().Version.AddVersion()
			//配置文件reload后 gc触发清理pool中的对象
			runtime.GC()
			log.Println("reload config version: ", GetCorePool().GetObj().Version.Version())
		}
	}()
	return &ConfRegister{Confer: &Conf{myConf}}

}

type ConfRegister struct {
	Confer types.Confer
}

func (c *ConfRegister) With(o ...types.Option) types.Register {
	return c
}

func (c *ConfRegister) Reload(ctx *types.Core) {
}

func (c *ConfRegister) Set(ctx *types.Core) {
	ctx.Config = c.Confer
}

func (c *ConfRegister) Unset(ctx *types.Core) {
}

//func (c *ConfReg) With(o ...types.Option) Register {
//	panic("implement me")
//}
//
//func (c *ConfReg) Reload(ctx *Core) {
//}
//
//func (c *ConfReg) Set(ctx *Core) {
//	ctx.Config = myConf
//}
//
//func (c *ConfReg) Unset(ctx *Core) {
//}
