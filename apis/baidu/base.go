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
	res, err := b.client.R("request root").Get("")
	return res, err
}

func (b *Api) GetRootNest() (*resty.Response, error) {
	res, err := b.client.R("request 1").Get("/1")
	return res, err
}
