package taobao

import (
	"github.com/fitan/gteml/pkg/httpclient"
	"github.com/fitan/gteml/pkg/types"
	"github.com/go-resty/resty/v2"
)

func NewTaoBaoApi(c *types.Context, client *resty.Client) *TaoBaoApi {
	return &TaoBaoApi{
		context: c,
		client:  httpclient.NewTraceClient(c.Tracer, client),
	}
}

type TaoBaoApi struct {
	context *types.Context
	client  *httpclient.TraceClient
}

func (t *TaoBaoApi) GetRoot() (*resty.Response, error) {
	return t.client.R().Get("/", "淘宝根目录")
}
