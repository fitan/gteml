package cache

import (
	"encoding/json"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

type Cache struct {
	ctx    *types.Context
	client *redis.Client
	option Option
}

type Option struct {
	Prefix string
}

func NewCache(ctx *types.Context, client *redis.Client, option Option) types.Cache {
	return &Cache{ctx: ctx, client: client, option: option}
}

func (c *Cache) genKey(key string) string {
	return c.option.Prefix + key
}

func (c *Cache) Get(key string, data interface{}) (string, bool, error) {
	key = c.genKey(key)
	log := c.ctx.Log.With(zap.String("funcName", "Get"), zap.String("redis key", key))
	val, err := c.client.Get(c.ctx.Tracer.SpanCtx("redis get "+key), key).Result()
	if err != nil {
		if err == redis.Nil {
			return val, false, nil
		}

		log.Error("redis get error", zap.Error(err))
		return val, false, err
	}
	if data != nil {
		json.Unmarshal([]byte(val), data)
	}
	return val, true, nil
}

func (c *Cache) GetCallBack(callBack func() (interface{}, error), key string, data interface{}, expiration time.Duration) (interface{}, error) {
	log := c.ctx.Log.With(zap.String("funcName", "GetCallBack"), zap.String("redis key", c.genKey(key)))
	_, has, err := c.Get(key, data)
	if err != nil {
		log.Error("redis err", zap.Error(err))
		return callBack()
	}
	if !has {
		c.ctx.Log.Warn("redis key is null", zap.Error(err))
		val, err := callBack()
		if err != nil {
			return val, err
		}

		_, valStr := json.Marshal(val)
		err = c.Put(key, valStr, expiration)
		if err != nil {
			log.Error("redis put error")
		}
		return val, err
	}

	return data, nil
}

func (c *Cache) Put(key string, val interface{}, expiration time.Duration) error {
	key = c.genKey(key)
	log := c.ctx.Log.With(zap.String("funcName", "Put"), zap.String("redis key", key))
	_, err := c.client.Set(c.ctx.Tracer.SpanCtx("redis put "+key), key, val, expiration).Result()
	if err != nil {
		log.Error("redis set error", zap.Error(err))
	}
	return err
}

func (c *Cache) PutCallBack(callBack func() (interface{}, error), key string) error {
	key = c.genKey(key)
	log := c.ctx.Log.With(zap.String("funcName", "PutCallBack"), zap.String("redis key", key))
	_, err := callBack()
	if err != nil {
		log.Error("callBack error", zap.Error(err))
		return err
	}

	_, err = c.Delete(key)
	if err != nil {
		c.ctx.Log.Error("redis delete error", zap.Error(err))
		return err
	}
	return nil
}

func (c *Cache) Delete(key string) (bool, error) {
	key = c.genKey(key)
	log := c.ctx.Log.With(zap.String("funcName", "Delete"), zap.String("redis key", key))
	_, err := c.client.Del(c.ctx.Tracer.SpanCtx("redis del "+key), key).Result()
	if err != nil {
		if err == redis.Nil {
			return true, nil
		}
		log.Error("redis delete error", zap.Error(err))
		return false, err
	}
	return true, nil
}

func (c *Cache) DeleteCallBack(callBack func() (interface{}, error), key string) error {
	key = c.genKey(key)
	log := c.ctx.Log.With(zap.String("funcName", "DeleteCallBack"), zap.String("redis key", key))

	_, err := callBack()
	if err != nil {
		log.Error("redis callBack", zap.Error(err))
		return err
	}

	_, err = c.Delete(key)
	if err != nil {
		log.Error("redis delete error", zap.Error(err))
		return err
	}

	return nil
}
