package core

import (
	"github.com/casbin/casbin/v2"
	log2 "github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/fitan/magic/dao"
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
	_ "github.com/go-sql-driver/mysql"
	mysqlapm "go.elastic.co/apm/module/apmgormv2/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type daoReg struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
}

func (s *daoReg) Reload(c *types.Core) {
	s.db = nil
}

func (s *daoReg) GetObj() *daoReg {
	if s.db == nil {
		db, err := gorm.Open(mysqlapm.Open(ConfReg.Confer.GetMyConf().Mysql.Url))
		if err != nil {
			log.Panicf("mysql create db: %s", err.Error())
		}

		s.db = db

		a, err := gormadapter.NewAdapterByDBWithCustomTable(db, &model.CasbinRule{})
		if err != nil {
			log.Panicf("casbin NewAdapterByDBWithCustomTable: %s", err.Error())
		}
		e, err := casbin.NewEnforcer("./rbac_model.conf", a, true)
		if err != nil {
			log.Panicf("casbin create enforcer: %s", err.Error())
		}
		logm := &log2.DefaultLogger{}
		logm.EnableLog(true)
		e.SetLogger(logm)
		if err != nil {
			log.Panicf("casbin create enforcer: %s", err.Error())
		}
		s.enforcer = e
		//s.enforcer.EnableEnforce(true)
		s.enforcer.AddNamedDomainMatchingFunc("g", "keyMatch", util.KeyMatch)
		//s.enforcer.SetRoleManager(s.enforcer.GetRoleManager())
	}
	return s
}

func (s *daoReg) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (s *daoReg) Set(c *types.Core) {
	obj := s.GetObj()
	c.Dao = dao.NewDAO(obj.db, obj.enforcer, c)
}

func (s *daoReg) Unset(c *types.Core) {

}
