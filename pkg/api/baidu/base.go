package baidu

import (
	"github.com/fitan/magic/pkg/httpclient"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

type BaiduApi struct {
	context *types.Context
	client  *httpclient.TraceClient
}

func NewBaiduApi(c *types.Context, client *resty.Client) *BaiduApi {
	return &BaiduApi{context: c, client: &httpclient.TraceClient{
		Tracer: c.Tracer,
		Client: client,
	}}
}

func (b *BaiduApi) GetSum() {
	b.GetRoot()
	b.GetRootNest()
}

func (b *BaiduApi) GetRoot() (*resty.Response, error) {
	res, err := b.client.R().Get("", "请求根目录")
	return res, err
}

func (b *BaiduApi) GetRootNest() (*resty.Response, error) {
	res, err := b.client.R().Get("/1", "请求根目录的子目录")
	return res, err
}
