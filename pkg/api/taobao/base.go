package taobao

import (
	"github.com/fitan/gteml/pkg/httpclient"
	"github.com/fitan/gteml/pkg/types"
	"github.com/go-resty/resty/v2"
)

var client = httpclient.NewClient(httpclient.WithHost("http://www.taobao.com"))

func NewTaoBaoApi(t *types.Context) *TaoBaoApi {
	return &TaoBaoApi{
		Context: t,
		TraceClient: &httpclient.TraceClient{
			Tracer: t.Tracer,
			Client: client,
		},
	}
}

type TaoBaoApi struct {
	Context *types.Context
	*httpclient.TraceClient
}

func (t *TaoBaoApi) GetRoot() (*resty.Response, error) {
	return t.R().Get("/", "淘宝根目录")
}
