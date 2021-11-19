package types

import (
	"context"
	"github.com/fitan/magic/model"
)

type DAOer interface {
	Storage() Storager
}

type Storager interface {
	User() UserModeler
}

type UserModeler interface {
	ById(ctx context.Context, id int64, preload ...string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	CheckPassword(ctx context.Context, userName string, password string) (*model.User, error)
}
