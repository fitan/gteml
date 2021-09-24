package core

//import (
//	"github.com/go-redis/redis/v8"
//	"go.uber.org/zap"
//)

//type Storage struct {
//	core *Context
//	DB *ent.Client
//	Redis *redis.Client
//}
//
//func (s *Storage) Get(callBack func() (interface{}, error),key string) (interface{}, error) {
//	val,err := s.Redis.Get(s.core.TraceLog.Context(), key).Result()
//	switch {
//	case err == redis.Nil:
//		s.core.TraceLog.Info("key does not exist", zap.String("key", key))
//	case err != nil:
//		s.core.TraceLog.Error("redis get key error", zap.Error(err), zap.String("key", key))
//	default:
//		return val,err
//	}
//	return callBack()
//}
//
//func (s *Storage) Update(callBack func() (interface{}, error),key string) (interface{}, error) {
//	val, err := callBack()
//	if err != nil {
//		s.core.TraceLog.Error("redis update callback error", zap.Error(err), zap.String("key", key))
//		return val, err
//	}
//
//	_, err = s.Redis.Del(s.core.TraceLog.Context(), key).Result()
//	switch {
//	case err == redis.Nil:
//		s.core.TraceLog.Info("redis update to delete key is nil")
//	case err != nil :
//		s.core.TraceLog.Error("redis update to delete key error", zap.Error(err), zap.String("key", key))
//	}
//	return val, nil
//
//
//}
//
//func (s *Storage) Delete(callBack func() (interface{}, error),key string) (interface{}, error) {
//	val, err := callBack()
//	if err != nil {
//		s.core.TraceLog.Error("redis delete error", zap.Error(err), zap.String("key", key))
//		return val, err
//	}
//	_,err = s.Redis.Del(s.core.TraceLog.Context(), key).Result()
//	switch  {
//	case err == redis.Nil:
//		s.core.TraceLog.Info("redis delete is nil")
//	case err != nil:
//		s.core.TraceLog.Error("redis delete error", zap.Error(err), zap.String("key", key))
//	}
//	return val, nil
//}
