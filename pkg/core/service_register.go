package core

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services"
)

type ServiceRegister struct {
}

func (s ServiceRegister) With(o ...types.Option) types.Register {
	return s
}

func (s ServiceRegister) Reload(c *types.Core) {
	return
}

func (s ServiceRegister) Set(c *types.Core) {
	c.Services = services.NewServices(c)
}

func (s ServiceRegister) Unset(c *types.Core) {
	return
}
