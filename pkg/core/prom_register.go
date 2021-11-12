package core

import (
	"github.com/fitan/magic/pkg/prometheus"
	"github.com/fitan/magic/pkg/types"
)

type PromRegister struct {
	prom types.Promer
}

func NewPromRegister() *PromRegister {
	return &PromRegister{prom: prometheus.NewGauge(prometheus.GetGinprom())}
}

func (p *PromRegister) Get() types.Promer {
	return p.prom
}

func (p *PromRegister) With(o ...types.Option) types.Register {
	return p
}

func (p *PromRegister) Reload(c *types.Core) {
}

func (p *PromRegister) Set(c *types.Core) {
	c.Prom = p.prom
}

func (p *PromRegister) Unset(c *types.Core) {
}
