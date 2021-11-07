package core

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/fitan/magic/ent"
	"github.com/fitan/magic/pkg/storage"
	"github.com/fitan/magic/pkg/types"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type storageReg struct {
	Client *ent.Client
}

func (s *storageReg) Reload(c *types.Core) {
	s.Client = nil
}

func (s *storageReg) GetClient(c *types.Core) *ent.Client {
	if s.Client == nil {
		db, err := sql.Open("mysql", c.Config.Mysql.Url)
		if err != nil {
			log.Panicf("mysql create db: %s", err.Error())
		}
		db.SetMaxIdleConns(Conf.MyConf.Mysql.MaxIdleConns)
		db.SetMaxOpenConns(Conf.MyConf.Mysql.MaxOpenConns)
		lt, err := time.ParseDuration(Conf.MyConf.Mysql.ConnMaxLifetime)
		if err != nil {
			log.Panicf("parse ConnMaxLifetime err: %s", err.Error())
		}
		it, err := time.ParseDuration(Conf.MyConf.Mysql.ConnMaxIdleTime)
		if err != nil {
			log.Panicf("parse ConnMaxIdleTime err: %s", err.Error())
		}
		db.SetConnMaxLifetime(lt)
		db.SetConnMaxIdleTime(it)
		s.Client = ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))

	}
	return s.Client
}

func (s *storageReg) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (s *storageReg) Set(c *types.Core) {
	c.Storage = storage.NewStorage(c, s.GetClient(c))
}

func (s *storageReg) Unset(c *types.Core) {

}
