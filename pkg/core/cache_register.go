package core

import (
	"github.com/fitan/magic/pkg/cache"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

type CacheReg struct {
	url    string
	client *redis.Client
}

func (c *CacheReg) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (c *CacheReg) Set(ctx *types.Context) {
	if ctx.Config.Redis.Url != c.url {
		client := redis.NewClient(&redis.Options{Addr: ctx.Config.Redis.Url, Password: "", DB: 0})
		client.AddHook(redisotel.NewTracingHook())
		c.url = ctx.Config.Redis.Url
		c.client = client
	}

	ctx.Cache = cache.NewCache(ctx, c.client, cache.Option{ctx.Config.App.Name, 5 * time.Second})
}

func (c *CacheReg) Unset(ctx *types.Context) {

}
