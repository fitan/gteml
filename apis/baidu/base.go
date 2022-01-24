package baidu

import (
	"github.com/fitan/magic/pkg/httpclient"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

type Api struct {
	context types.ServiceCore
	client  *httpclient.TraceClient
}

func NewApi(c types.ServiceCore, client *resty.Client) *Api {
	return &Api{context: c, client: &httpclient.TraceClient{
		Tracer: c.GetTrace(),
		Client: client,
	}}
}

func (b *Api) GetSum() {
	b.GetRoot()
	b.GetRootNest()
}

func (b *Api) GetRoot() (*resty.Response, error) {
	res, err := b.client.R().Get("", "请求根目录")
	return res, err
}

func (b *Api) GetRootNest() (*resty.Response, error) {
	res, err := b.client.R().Get("/1", "请求根目录的子目录")
	return res, err
}
