package core

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services"
)

type ServiceRegister struct {
	enforver *casbin.Enforcer
}

func (s *ServiceRegister) GetEnforcer() *casbin.Enforcer {
	//if s.enforver == nil {
	//	a, err := gormadapter.NewAdapter("mysql", ConfReg.Confer.GetMyConf().Mysql.Url)
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//	e, err := casbin.NewEnforcer("./rbac_model.conf", a)
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//
	//	s.enforver = e
	//}
	return s.enforver
}

func (s *ServiceRegister) With(o ...types.Option) types.Register {
	return s
}

func (s *ServiceRegister) Reload(c *types.Core) {
	s.enforver = nil
}

func (s *ServiceRegister) Set(c *types.Core) {
	c.Services = services.NewServices(c, s.GetEnforcer())
}

func (s *ServiceRegister) Unset(c *types.Core) {
	return
}
