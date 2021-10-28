package core

import (
	"github.com/fitan/magic/pkg/ent"
	"github.com/fitan/magic/pkg/storage"
	"github.com/fitan/magic/pkg/types"
	"log"
)

type storageReg struct {
	Client *ent.Client
}

func (s *storageReg) Reload(c *types.Context) {
	s.Client = nil
}

func (s *storageReg) GetClient(c *types.Context) *ent.Client {
	if s.Client == nil {
		client, err := ent.Open("mysql", c.Config.Mysql.Url)
		if err != nil {
			log.Println("mysql create client: ", err.Error())
		} else {
			s.Client = client
		}
	}
	return s.Client
}

func (s *storageReg) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (s *storageReg) Set(c *types.Context) {
	c.Storage = storage.NewStorage(c, s.Client)
}

func (s *storageReg) Unset(c *types.Context) {

}
