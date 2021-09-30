package core

import (
	"github.com/fitan/gteml/pkg/api/baidu"
	"github.com/fitan/gteml/pkg/api/taobao"
)

type Apis struct {
	Baidu  *baidu.BaiduApi
	Taobao *taobao.TaoBaoApi
}

type ApisRegister struct {
}

func (h *ApisRegister) With(o ...Option) Register {
	return h
}

func (h *ApisRegister) Set(c *Context) {
	c.Apis.Baidu = baidu.NewBaiduApi(c)
	c.Apis.Taobao = taobao.NewTaoBaoApi(c)
}

func (h *ApisRegister) Unset(c *Context) {
	c.Apis.Baidu = nil
	c.Apis.Taobao = nil
}
