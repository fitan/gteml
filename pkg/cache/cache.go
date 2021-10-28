package cache

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type Cache struct {
	core   *types.Context
	client *redis.Client
	option Option
}

type Option struct {
	Prefix  string
	SetTime time.Duration
}

func NewCache(core *types.Context, client *redis.Client, option Option) types.Cache {
	return &Cache{core: core, client: client, option: option}
}

func (c *Cache) genKey(objStr string, id int) string {
	return c.option.Prefix + "." + objStr + "." + strconv.Itoa(id)
}

func (c *Cache) Get(objStr string, id int) (interface{}, bool, error) {
	log := c.core.Log.With(zap.String("funcName", "Get"), zap.String("redis key", c.genKey(objStr, id)))
	key := c.genKey(objStr, id)
	val, err := c.client.Get(c.core.Tracer.SpanCtx("redis get "+key), key).Result()
	if err != nil {
		if err == redis.Nil {
			return val, false, nil
		}

		log.Error("redis get error", zap.Error(err))
		return val, false, err
	}
	return val, true, nil
}

func (c *Cache) GetCallBack(callBack func() (interface{}, error), objStr string, id int) (interface{}, error) {
	log := c.core.Log.With(zap.String("funcName", "GetCallBack"), zap.String("redis key", c.genKey(objStr, id)))
	val, has, err := c.Get(objStr, id)
	if err != nil {
		log.Error("redis err", zap.Error(err))
		return callBack()
	}
	if !has {
		c.core.Log.Info("redis key is null", zap.Error(err))
		val, err := callBack()
		if err != nil {
			return val, err
		}

		err = c.Put(objStr, id, val)
		if err != nil {
			log.Error("redis put error")
		}
		return val, err
	}
	return val, nil
}

func (c *Cache) Put(objStr string, id int, val interface{}) error {
	log := c.core.Log.With(zap.String("funcName", "Put"), zap.String("redis key", c.genKey(objStr, id)))
	key := c.genKey(objStr, id)
	_, err := c.client.Set(c.core.Tracer.SpanCtx("redis put "+key), key, val, c.option.SetTime).Result()
	if err != nil {
		log.Error("redis set error", zap.Error(err))
	}
	return err
}

func (c *Cache) PutCallBack(callBack func() (interface{}, error), objStr string, id int) error {
	log := c.core.Log.With(zap.String("funcName", "PutCallBack"), zap.String("redis key", c.genKey(objStr, id)))
	_, err := callBack()
	if err != nil {
		log.Error("callBack error", zap.Error(err))
		return err
	}

	_, err = c.Delete(objStr, id)
	if err != nil {
		c.core.Log.Error("redis delete error", zap.Error(err))
		return err
	}
	return nil
}

func (c *Cache) Delete(objStr string, id int) (bool, error) {
	log := c.core.Log.With(zap.String("funcName", "Delete"), zap.String("redis key", c.genKey(objStr, id)))
	key := c.genKey(objStr, id)
	_, err := c.client.Del(c.core.Tracer.SpanCtx("redis del "+key), key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		log.Error("redis delete error", zap.Error(err))
		return false, err
	}
	return true, nil
}

func (c *Cache) DeleteCallBack(callBack func() (interface{}, error), objStr string, id int) error {
	log := c.core.Log.With(zap.String("funcName", "DeleteCallBack"), zap.String("redis key", c.genKey(objStr, id)))
	_, err := callBack()
	if err != nil {
		log.Error("redis callBack", zap.Error(err))
		return err
	}

	_, err = c.Delete(objStr, id)
	if err != nil {
		log.Error("redis delete error", zap.Error(err))
		return err
	}
	return nil
}
