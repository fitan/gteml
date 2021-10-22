package core

import (
	"github.com/fitan/gteml/pkg/ent"
	"github.com/fitan/gteml/pkg/ent/user"
	"github.com/fitan/gteml/pkg/types"
	"log"
)

type Storage struct {
	core *types.Context
	*ent.Client
}

func (s *Storage) GetById(id int) (*ent.User, error) {
	return s.Client.User.Get(s.core.Tracer.SpanCtx("GetById"), id)
}

func (s *Storage) GetByIds(ids []int) ([]*ent.User, error) {
	return s.Client.User.Query().Where(user.IDIn(ids...)).All(s.core.Tracer.SpanCtx("GetByIds"))
}

func NewStorage(core *types.Context, client *ent.Client) types.Storage {
	return &Storage{
		core:   core,
		Client: client,
	}
}

type StorageReg struct {
	url    string
	client *ent.Client
}

func (s StorageReg) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (s StorageReg) Set(c *types.Context) {
	if c.Config.Mysql.Url != s.url {
		s.url = c.Config.Mysql.Url
		client, err := ent.Open("mysql", c.Config.Mysql.Url)
		if err != nil {
			log.Println("ent mysql: ", err.Error())
		}
		s.client = client
	}

	c.Storage = NewStorage(c, s.client)
}

func (s StorageReg) Unset(c *types.Context) {

}
