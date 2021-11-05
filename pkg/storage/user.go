package storage

import (
	ent "github.com/fitan/magic/ent"
	"github.com/fitan/magic/ent/user"
	"github.com/fitan/magic/pkg/types"
)

type User struct {
	ctx    *types.Context
	client *ent.Client
}

func NewUser(ctx *types.Context, client *ent.Client) *User {
	return &User{ctx: ctx, client: client}
}

func (s *User) GetById(id int) (res *ent.User, err error) {
	return s.client.User.Get(s.ctx.Tracer.SpanCtx("GetById"), id)
}

func (s *User) GetByIds(ids []int) (res []*ent.User, err error) {
	return s.client.User.Query().Where(user.IDIn(ids...)).All(s.ctx.Tracer.SpanCtx("GetByIds"))
}
