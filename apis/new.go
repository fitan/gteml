package apis

import (
	"github.com/fitan/magic/apis/baidu"
	"github.com/fitan/magic/apis/taobao"
	apisTypes "github.com/fitan/magic/apis/types"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

type Apis struct {
	BaiduApi  *baidu.BaiduApi
	TaobaoApi *taobao.TaoBaoApi
}

func NewApis(core types.ServiceCore, baiduApi, taobaoApi *resty.Client) apisTypes.Apis {
	return &Apis{
		BaiduApi:  baidu.NewBaiduApi(core, baiduApi),
		TaobaoApi: taobao.NewTaoBaoApi(core, taobaoApi),
	}
}

func (a *Apis) Baidu() apisTypes.BaiduApi {
	return a.BaiduApi
}

func (a *Apis) Taobao() apisTypes.TaobaoApi {
	return a.TaobaoApi
}
