package storage

import (
	ent "github.com/fitan/magic/pkg/ent"
	"github.com/fitan/magic/pkg/ent/user"
	"github.com/fitan/magic/pkg/types"
)

type User struct {
	ctx       *types.Context
	client    *ent.Client
	openCache bool
}

func NewUser(ctx *types.Context, client *ent.Client, openCache bool) *User {
	return &User{ctx: ctx, client: client, openCache: openCache}
}

func (s *User) GetById(id int) (res *ent.User, err error) {
	return s.client.User.Get(s.ctx.Tracer.SpanCtx("GetById"), id)
}

func (s *User) GetByIds(ids []int) ([]*ent.User, error) {
	return s.client.User.Query().Where(user.IDIn(ids...)).All(s.ctx.Tracer.SpanCtx("GetByIds"))
}
