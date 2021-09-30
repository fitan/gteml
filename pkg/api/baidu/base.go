package baidu

import (
	"github.com/fitan/gteml/pkg/common"
	"github.com/fitan/gteml/pkg/httpclient"
	"github.com/fitan/gteml/pkg/trace"
	"github.com/go-resty/resty/v2"
)

var client = httpclient.NewClient(httpclient.WithHost("http://www.baidu.c"), httpclient.WithTrace(trace.GetTp(), "baidu", false))

type BaiduApi struct {
	Context common.Context
	client  *httpclient.TraceClient
}

func NewBaiduApi(t common.Context) *BaiduApi {
	return &BaiduApi{Context: t, client: &httpclient.TraceClient{
		Tracer: t,
		Client: client,
	}}
}

func (b *BaiduApi) GetRoot() (*resty.Response, error) {
	res, err := b.client.R().Get("/fsfds", "请求根目录")
	return res, err
}
