package core

import (
	"github.com/fitan/gteml/pkg/types"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type Cache struct {
	core   *types.Context
	Client *redis.Client
	option option
}

type option struct {
	prefix  string
	setTime time.Duration
}

func NewCache(core *types.Context, client *redis.Client, option option) types.Cache {
	return &Cache{core: core, Client: client, option: option}
}

func (c *Cache) genKey(objStr string, id int) string {
	return c.option.prefix + "." + objStr + "." + strconv.Itoa(id)
}

func (c *Cache) Get(objStr string, id int) (interface{}, bool, error) {
	key := c.genKey(objStr, id)
	val, err := c.Client.Get(c.core.Tracer.SpanCtx("redis get "+key), key).Result()
	if err != nil {
		if err == redis.Nil {
			return val, false, nil
		}

		return val, false, err
	}
	return val, true, nil
}

func (c *Cache) GetCallBack(callBack func() (interface{}, error), objStr string, id int) (interface{}, error) {
	val, has, err := c.Get(objStr, id)
	if err != nil {
		c.core.Log.Error("redis getCallback key error", zap.Error(err), zap.String("redis key", c.genKey(objStr, id)))
		return callBack()
	}
	if !has {
		c.core.Log.Info("redis getCallBack key is null", zap.String("redis key", c.genKey(objStr, id)))
		val, err := callBack()
		if err != nil {
			return val, err
		}

		err = c.Put(objStr, id, val)
		if err != nil {
			c.core.Log.Error("redis getCallback put val error", zap.String("key", c.genKey(objStr, id)))
		}
		return val, err
	}
	return val, nil
}

func (c *Cache) Put(objStr string, id int, val interface{}) error {
	key := c.genKey(objStr, id)
	_, err := c.Client.Set(c.core.Tracer.SpanCtx("redis put "+key), key, val, c.option.setTime).Result()
	if err != nil {
		c.core.Log.Error("redis put")
	}
	return err
}

func (c *Cache) PutCallBack(callBack func() (interface{}, error), objStr string, id int) error {
	_, err := callBack()
	if err != nil {
		c.core.Log.Error("putCallback callback error", zap.Error(err))
		return err
	}

	_, err = c.Delete(objStr, id)
	if err != nil {
		c.core.Log.Error("putCallback redis del error", zap.Error(err), zap.String("key", c.genKey(objStr, id)))
		return err
	}
	return nil
}

func (c *Cache) Delete(objStr string, id int) (bool, error) {
	key := c.genKey(objStr, id)
	_, err := c.Client.Del(c.core.Tracer.SpanCtx("redis del "+key), key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (c *Cache) DeleteCallBack(callBack func() (interface{}, error), objStr string, id int) error {
	_, err := callBack()
	if err != nil {
		c.core.Log.Error("redis deleteCallback callback error", zap.Error(err))
		return err
	}

	_, err = c.Delete(objStr, id)
	if err != nil {
		c.core.Log.Error("redis deleteCallback del key error", zap.Error(err), zap.String("key", c.genKey(objStr, id)))
		return err
	}
	return nil
}

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
		c.url = ctx.Config.Redis.Url
		c.client = client
	}

	ctx.Cache = NewCache(ctx, c.client, option{ctx.Config.App.Name, 5 * time.Second})
}

func (c *CacheReg) Unset(ctx *types.Context) {

}
