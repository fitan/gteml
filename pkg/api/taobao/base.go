package taobao

import (
	"github.com/fitan/gteml/pkg/common"
	"github.com/fitan/gteml/pkg/httpclient"
	"github.com/go-resty/resty/v2"
)

var client = httpclient.NewClient(httpclient.WithHost("http://www.taobao.com"))

func NewTaoBaoApi(t common.Context) *TaoBaoApi {
	return &TaoBaoApi{
		Context: t,
		TraceClient: &httpclient.TraceClient{
			Tracer: t,
			Client: client,
		},
	}
}

type TaoBaoApi struct {
	Context common.Context
	*httpclient.TraceClient
}

func (t *TaoBaoApi) GetRoot() (*resty.Response, error) {
	return t.R().Get("/", "淘宝根目录")
}
