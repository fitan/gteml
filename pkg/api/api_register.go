package api

import (
	"github.com/fitan/gteml/pkg/api/baidu"
	"github.com/fitan/gteml/pkg/api/taobao"
	"github.com/fitan/gteml/pkg/httpclient"
	"github.com/fitan/gteml/pkg/trace"
	"github.com/fitan/gteml/pkg/types"
	"github.com/go-resty/resty/v2"
	"sync"
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
	m            sync.Mutex
}

func (h *ApisRegister) getApis(c *types.Context) *Apis {
	h.m.Lock()
	defer h.m.Unlock()

	if h.baiduClient == nil {
		h.baiduClient = httpclient.NewClient(
			httpclient.WithHost(c.Config.Api.Baidu.Url),
			httpclient.WithTrace(trace.GetTp(), "baidu", c.Config.Api.Baidu.TraceDebug),
			httpclient.WithDebug(c.Config.Api.Baidu.RestyDebug))
	}

	if h.taobaoClient == nil {
		h.taobaoClient = httpclient.NewClient(
			httpclient.WithHost(c.Config.Api.Taobao.Url),
			httpclient.WithTrace(trace.GetTp(), "taobao", c.Config.Api.Taobao.TraceDebug),
			httpclient.WithDebug(c.Config.Api.Taobao.RestyDebug))
	}

	return &Apis{
		BaiduApi:  baidu.NewBaiduApi(c, h.baiduClient),
		TaobaoApi: taobao.NewTaoBaoApi(c, h.taobaoClient),
	}
}

func (h *ApisRegister) Reload(c *types.Context) {
	h.m.Lock()
	defer h.m.Unlock()
	h.taobaoClient = nil
	h.baiduClient = nil
}

func (h *ApisRegister) With(o ...types.Option) types.Register {
	return h
}

func (h *ApisRegister) Set(c *types.Context) {
	c.Apis = h.getApis(c)
}

func (h *ApisRegister) Unset(c *types.Context) {
}
