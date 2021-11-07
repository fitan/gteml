package api

import (
	"github.com/fitan/magic/pkg/api/baidu"
	"github.com/fitan/magic/pkg/api/taobao"
	"github.com/fitan/magic/pkg/httpclient"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

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
	baiduClient  *resty.Client
	taobaoClient *resty.Client
}

func (h *ApisRegister) getApis(c *types.Core) *Apis {
	if h.baiduClient == nil {
		h.baiduClient = httpclient.NewClient(
			httpclient.WithHost(c.Config.Api.Baidu.Url),
			httpclient.WithTrace(c.Tracer.Tp(), "baidu", c.Config.Api.Baidu.TraceDebug),
			httpclient.WithDebug(c.Config.Api.Baidu.RestyDebug))
	}

	if h.taobaoClient == nil {
		h.taobaoClient = httpclient.NewClient(
			httpclient.WithHost(c.Config.Api.Taobao.Url),
			httpclient.WithTrace(c.Tracer.Tp(), "taobao", c.Config.Api.Taobao.TraceDebug),
			httpclient.WithDebug(c.Config.Api.Taobao.RestyDebug))
	}

	return &Apis{
		BaiduApi:  baidu.NewBaiduApi(c, h.baiduClient),
		TaobaoApi: taobao.NewTaoBaoApi(c, h.taobaoClient),
	}
}

func (h *ApisRegister) Reload(c *types.Core) {
	h.taobaoClient = nil
	h.baiduClient = nil
}

func (h *ApisRegister) With(o ...types.Option) types.Register {
	return h
}

func (h *ApisRegister) Set(c *types.Core) {
	c.Apis = h.getApis(c)
}

func (h *ApisRegister) Unset(c *types.Core) {
}
