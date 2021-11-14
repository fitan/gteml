package core

import (
	"github.com/fitan/magic/pkg/cache"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

type CacheReg struct {
	client *redis.Client
}

func (c *CacheReg) Reload(ctx *types.Core) {
	c.client = nil
}

func (c *CacheReg) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (c *CacheReg) Set(ctx *types.Core) {
	if c.client == nil {
		c.client = redis.NewClient(&redis.Options{Addr: ctx.Config.GetMyConf().Redis.Url, Password: ctx.Config.GetMyConf().Redis.Password, DB: ctx.Config.GetMyConf().Redis.Db})
		c.client.AddHook(redisotel.NewTracingHook())
	}

	ctx.Cache = cache.NewCache(ctx, c.client, cache.Option{Prefix: ctx.Config.GetMyConf().App.Name})
}

func (c *CacheReg) Unset(ctx *types.Core) {

}
