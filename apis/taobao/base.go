package taobao

import (
	"github.com/fitan/magic/pkg/httpclient"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

func NewTaoBaoApi(c types.ServiceCore, client *resty.Client) *Api {
	return &Api{
		context: c,
		client:  httpclient.NewTraceClient(c.GetTrace(), client),
	}
}

type Api struct {
	context types.ServiceCore
	client  *httpclient.TraceClient
}

func (t *Api) GetRoot() (*resty.Response, error) {
	return t.client.R("request taobao root").Get("/")
}
