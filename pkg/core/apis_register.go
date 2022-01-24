package core

import (
	"github.com/fitan/magic/apis"
	apiTypes "github.com/fitan/magic/apis/types"
	"github.com/fitan/magic/pkg/httpclient"
	"github.com/fitan/magic/pkg/micro"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

type ApisRegister struct {
	baiduClient  *resty.Client
	taobaoClient *resty.Client
	gtemlClient  *resty.Client
}

func (h *ApisRegister) getApis(c *types.Core) apiTypes.Apis {
	if h.baiduClient == nil {
		h.baiduClient = httpclient.NewClient(
			httpclient.WithHost(c.GetConfig().GetMyConf().Apis.Baidu.Url),
			httpclient.WithTrace(c.Tracer.Tp(), c.GetConfig().GetMyConf().Apis.Baidu.TraceDebug),
			httpclient.WithDebug(c.GetConfig().GetMyConf().Apis.Baidu.RestyDebug))
	}

	if h.taobaoClient == nil {
		h.taobaoClient = httpclient.NewClient(
			httpclient.WithHost(c.GetConfig().GetMyConf().Apis.Taobao.Url),
			httpclient.WithTrace(c.Tracer.Tp(), c.GetConfig().GetMyConf().Apis.Taobao.TraceDebug),
			httpclient.WithDebug(c.GetConfig().GetMyConf().Apis.Taobao.RestyDebug))
	}

	if h.gtemlClient == nil {
		h.gtemlClient = httpclient.NewClient(
			httpclient.WithMicroHost("gteml", micro.GetConsulRegistry()),
			httpclient.WithTrace(c.Tracer.Tp(), c.GetConfig().GetMyConf().Apis.Taobao.TraceDebug),
			httpclient.WithDebug(c.GetConfig().GetMyConf().Apis.Taobao.RestyDebug))
	}

	return apis.NewApis(c, h.baiduClient, h.taobaoClient, h.gtemlClient)
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
