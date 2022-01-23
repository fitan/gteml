package taobao

import (
	"github.com/fitan/magic/pkg/httpclient"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

func NewTaoBaoApi(c types.ServiceCore, client *resty.Client) *TaoBaoApi {
	return &TaoBaoApi{
		context: c,
		client:  httpclient.NewTraceClient(c.GetTrace(), client),
	}
}

type TaoBaoApi struct {
	context types.ServiceCore
	client  *httpclient.TraceClient
}

func (t *TaoBaoApi) GetRoot() (*resty.Response, error) {
	return t.client.R().Get("/", "淘宝根目录")
}
