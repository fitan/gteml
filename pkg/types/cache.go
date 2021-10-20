package types

type Cache interface {
	Get(objStr string, id int) (interface{}, bool, error)
	Put(objStr string, id int, val interface{}) error
	Delete(objStr string, id int) (bool, error)
	GetCallBack(callBack func() (interface{}, error), objStr string, id int) (interface{}, error)
	PutCallBack(callBack func() (interface{}, error), objStr string, id int) error
	DeleteCallBack(callBack func() (interface{}, error), objStr string, id int) error
}
