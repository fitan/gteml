package storage

import (
	"context"
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
	"gorm.io/gorm"
)

func NewUser(db *gorm.DB) types.UserModeler {
	return &User{db: db}
}

type User struct {
	db *gorm.DB
}

func (u *User) CheckPassword(ctx context.Context, userName string, password string) (*model.User, error) {
	db := u.db
	res := &model.User{}
	err := db.WithContext(ctx).Where("email = ? And password = ?", userName, password).First(res).Error
	return res, err
}

func (u *User) ById(ctx context.Context, id int64, preload ...string) (*model.User, error) {
	db := u.db.WithContext(ctx)
	if len(preload) != 0 {
		for _, v := range preload {
			db.Preload(v)
		}
	}

	res := model.User{}
	return &res, db.Where("fasdf").First(&res, id).Error
}

func (u *User) Create(user *model.User) error {
	db := u.db

	return db.Create(user).Error
}

func (u *User) Update(user *model.User) error {
	db := u.db

	return db.Save(user).Error
}
