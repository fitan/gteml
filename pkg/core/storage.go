package core

import (
	"github.com/fitan/gteml/pkg/ent"
	"github.com/fitan/gteml/pkg/ent/user"
)

type Storage struct {
	core *Context
	ent.Client
}

func (s *Storage) GetById(id int) (*ent.User, error) {
	return s.Client.User.Get(s.core.SpanCtx("GetById"), id)
}

func (s *Storage) GetByIds(ids []int) ([]*ent.User, error) {
	return s.Client.User.Query().Where(user.IDIn(ids...)).All(s.core.SpanCtx("GetByIds"))
}
