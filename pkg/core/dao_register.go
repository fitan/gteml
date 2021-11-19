package core

import (
	"github.com/fitan/magic/dao"
	"github.com/fitan/magic/pkg/types"
	_ "github.com/go-sql-driver/mysql"
	mysqlapm "go.elastic.co/apm/module/apmgormv2/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type daoReg struct {
	dao types.DAOer
}

func (s *daoReg) Reload(c *types.Core) {
	s.dao = nil
}

func (s *daoReg) GetDao(c *types.Core) types.DAOer {
	if s.dao == nil {
		db, err := gorm.Open(mysqlapm.Open(ConfReg.Confer.GetMyConf().Mysql.Url))
		if err != nil {
			log.Panicf("mysql create db: %s", err.Error())
		}

		s.dao = dao.NewDAO(db)
	}
	return s.dao
}

func (s *daoReg) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (s *daoReg) Set(c *types.Core) {
	c.Dao = s.GetDao(c)
}

func (s *daoReg) Unset(c *types.Core) {

}
