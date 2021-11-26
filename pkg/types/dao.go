package types

import (
	"gorm.io/gorm"
)

type DAOer interface {
	Storage() Storager
}

type Storager interface {
	User() User
	Role() Role
	Permission() Permission
	DB() *gorm.DB
}
