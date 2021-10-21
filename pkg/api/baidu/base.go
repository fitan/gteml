package baidu

import (
	"github.com/fitan/gteml/pkg/httpclient"
	"github.com/fitan/gteml/pkg/trace"
	"github.com/fitan/gteml/pkg/types"
	"github.com/go-resty/resty/v2"
)

var client = httpclient.NewClient(httpclient.WithHost("http://www.baidu.com"), httpclient.WithTrace(trace.GetTp(), "baidu", false))

type BaiduApi struct {
	Context *types.Context
	client  *httpclient.TraceClient
}

func NewBaiduApi(t *types.Context) *BaiduApi {
	return &BaiduApi{Context: t, client: &httpclient.TraceClient{
		Tracer: t.Tracer,
		Client: client,
	}}
}

func (b *BaiduApi) GetRoot() (*resty.Response, error) {
	res, err := b.client.R().Get("", "请求根目录")
	return res, err
}
