package types

import "github.com/fitan/magic/pkg/ent"

type Storage interface {
	User() UserInterface
}

type UserInterface interface {
	GetById(id int) (*ent.User, error)
	GetByIds(ids []int) ([]*ent.User, error)
}
