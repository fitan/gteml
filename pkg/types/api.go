package types

import "github.com/go-resty/resty/v2"

type Apis interface {
	Baidu() BaiduApi
	Taobao() TaobaoApi
}

type BaiduApi interface {
	GetRoot() (*resty.Response, error)
}

type TaobaoApi interface {
	GetRoot() (*resty.Response, error)
}
