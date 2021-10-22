package core

import (
	"github.com/fitan/gteml/pkg/api/baidu"
	"github.com/fitan/gteml/pkg/api/taobao"
	"github.com/fitan/gteml/pkg/types"
)

func NewApis(c *types.Context) *Apis {
	return &Apis{
		BaiduApi:  baidu.NewBaiduApi(c),
		TaobaoApi: taobao.NewTaoBaoApi(c),
	}
}

type Apis struct {
	BaiduApi  *baidu.BaiduApi
	TaobaoApi *taobao.TaoBaoApi
}

func (a *Apis) Baidu() types.BaiduApi {
	return a.BaiduApi
}

func (a *Apis) Taobao() types.TaobaoApi {
	return a.TaobaoApi
}

type ApisRegister struct {
}

func (h *ApisRegister) Reload(c *types.Context) {

}

func (h *ApisRegister) With(o ...types.Option) types.Register {
	return h
}

func (h *ApisRegister) Set(c *types.Context) {
	c.Apis = NewApis(c)
}

func (h *ApisRegister) Unset(c *types.Context) {
	//c.Apis.Baidu = nil
	//c.Apis.Taobao = nil
}
