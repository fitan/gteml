package core

import (
	"github.com/casbin/casbin/v2"
	log2 "github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/fitan/magic/dao"
	"github.com/fitan/magic/dao/dal/model"
	query2 "github.com/fitan/magic/dao/dal/query"
	"github.com/fitan/magic/pkg/dbquery"
	"github.com/fitan/magic/pkg/types"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DaoRegister struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
	query    *query2.Query
}

func (s *DaoRegister) Reload(c *types.Core) {
	s.db = nil
}

func (s *DaoRegister) GetObj(c *types.Core) *DaoRegister {
	if s.db == nil {
		db, err := gorm.Open(mysql.Open(ConfReg.Confer.GetMyConf().Mysql.Url), &gorm.Config{})
		//db, err := gorm.Open(mysqlapm.Open(ConfReg.Confer.GetMyConf().Mysql.Url))
		if err != nil {
			log.Panicf("mysql create db: %s", err.Error())
		}
		db.Use(otelgorm.NewPlugin())

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

		s.query = query2.Use(db)
		//s.enforcer.SetRoleManager(s.enforcer.GetRoleManager())
	}
	return s
}

func (s *DaoRegister) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (s *DaoRegister) Set(c *types.Core) {
	obj := s.GetObj(c)
	wrapQuery := &dbquery.WrapQuery{c, obj.query}
	c.Dao = dao.NewDAO(obj.db, wrapQuery, obj.enforcer, c)
}

func (s *DaoRegister) Unset(c *types.Core) {

}
