package types

import "github.com/go-resty/resty/v2"

type Apis interface {
	Baidu() BaiduApi
	Taobao() TaobaoApi
	Gteml() GtemlApi
}

type BaiduApi interface {
	GetRoot() (*resty.Response, error)
	GetRootNest() (*resty.Response, error)
	GetSum()
}

type TaobaoApi interface {
	GetRoot() (*resty.Response, error)
}

type GtemlApi interface {
	GetRoot(token string) (*resty.Response, error)
}
