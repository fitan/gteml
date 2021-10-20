package types

import "github.com/fitan/gteml/pkg/ent"

type Storage interface {
	GetById(id int) (*ent.User, error)
	GetByIds(ids []int) ([]*ent.User, error)
}
