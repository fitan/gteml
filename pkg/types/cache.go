package types

import "time"

type Cache interface {
	Get(key string, data interface{}) (string, bool, error)
	GetCallBack(
		callBack func() (interface{}, error), key string, data interface{}, expiration time.Duration,
	) (interface{}, error)
	Put(key string, val interface{}, expiration time.Duration) error
	PutCallBack(callBack func() (interface{}, error), key string) error
	Delete(key string) (bool, error)
	DeleteCallBack(callBack func() (interface{}, error), key string) error
}
