package gteml

import (
	"github.com/fitan/magic/pkg/httpclient"
	"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
)

func NewApi(c types.ServiceCore, client *resty.Client) *Api {
	return &Api{
		context: c,
		client:  httpclient.NewTraceClient(c.GetTrace(), client.SetDebug(true)),
	}
}

type Api struct {
	context types.ServiceCore
	client  *httpclient.TraceClient
}

func (t *Api) GetRoot(token string) (*resty.Response, error) {
	return t.client.R("getroot").SetHeader("Authorization", token).Get("/say")
}
