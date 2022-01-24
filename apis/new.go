package apis

import (
	"github.com/fitan/magic/apis/baidu"
	"github.com/fitan/magic/apis/gteml"
	"github.com/fitan/magic/apis/taobao"
	apisTypes "github.com/fitan/magic/apis/types"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

type Apis struct {
	BaiduApi  *baidu.Api
	TaobaoApi *taobao.Api
	GtemlApi  *gteml.Api
}

func NewApis(core types.ServiceCore, baiduApi, taobaoApi, gtemlApi *resty.Client) apisTypes.Apis {
	return &Apis{
		BaiduApi:  baidu.NewApi(core, baiduApi),
		TaobaoApi: taobao.NewTaoBaoApi(core, taobaoApi),
		GtemlApi:  gteml.NewApi(core, gtemlApi),
	}
}
func (a *Apis) Gteml() apisTypes.GtemlApi {
	return a.GtemlApi
}

func (a *Apis) Baidu() apisTypes.BaiduApi {
	return a.BaiduApi
}

func (a *Apis) Taobao() apisTypes.TaobaoApi {
	return a.TaobaoApi
}
