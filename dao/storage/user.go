package storage

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/model"
	"github.com/fitan/magic/pkg/types"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"path"
	"path/filepath"
	"strconv"
)

func NewUser(daoCore types.DaoCore, enforcer *casbin.Enforcer) types.User {
	return &User{core: daoCore, enforcer: enforcer}
}

type User struct {
	core     types.DaoCore
	enforcer *casbin.Enforcer
}

func (u *User) CheckUserPermission(userID uint, serviceID uint, path, method string) (err error) {
	fmt.Println(userID, serviceID, path, method)
	db := u.core.GetDao().Storage().DB()
	service := &model.Service{}
	err = db.First(service, serviceID).Error
	if err != nil {
		return errors.WithMessage(err, "service not found")
	}

	var domain string
	if service.ParentId != 0 {
		domain = filepath.Join(strconv.Itoa(int(service.ParentId)), strconv.Itoa(int(service.ID)))
	} else {
		domain = filepath.Join(strconv.Itoa(int(service.ID)))
	}

	user := userID2CasbinKey(userID)
	ok, err := u.enforcer.Enforce(user, domain, path, method)
	//ok, err := u.enforcer.Enforce(strconv.Itoa(int(userID)), domain, path, method)
	fmt.Println(user, domain, path, method, ok, err)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("permission denied")
	}
	return nil
}

func (u *User) CheckPassword(userName string, password string) (*model.User, error) {
	db := u.core.GetDao().Storage().DB()
	res := &model.User{}
	err := db.Where("email = ? AND pass_word = ?", userName, password).First(res).Error
	return res, err
}

func (u *User) ById(id uint, preload ...string) (*model.User, error) {
	db := u.core.GetDao().Storage().DB()
	if len(preload) != 0 {
		for _, v := range preload {
			db.Preload(v)
		}
	}

	res := model.User{}
	return &res, db.First(&res, id).Error
}

func (u *User) Create(user *model.User) error {
	db := u.core.GetDao().Storage().DB()

	return db.Create(user).Error
}

func (u *User) Update(user *model.User) error {
	db := u.core.GetDao().Storage().DB()

	return db.Save(user).Error
}

func (u *User) UnBindPermission(userID, roleID, serviceID uint) (err error) {
	db := u.core.GetDao().Storage().DB()

	var domain string
	var serviceType string

	err = db.Transaction(
		func(tx *gorm.DB) error {
			service := &model.Service{}
			err = db.First(service, serviceID).Error
			if err != nil {
				return err
			}

			if service.ParentId != 0 {
				serviceType = "service"
				domain = filepath.Join(strconv.Itoa(int(service.ParentId)), strconv.Itoa(int(service.ID)))
			} else {
				serviceType = "project"
				domain = filepath.Join(strconv.Itoa(int(service.ID)), "*")
			}

			userStr := userID2CasbinKey(userID)
			roleStr := roleID2CasbinKey(roleID)

			if serviceType == "project" {
				filterPolicies := make([][]string, 0, 0)
				policies := u.enforcer.GetFilteredGroupingPolicy(0, userStr, roleStr)
				for _, v := range policies {
					has, err := filepath.Match(domain, v[2])
					if err != nil {
						return err
					}

					if has {
						filterPolicies = append(filterPolicies, v)
					}
				}

				_, err = u.enforcer.RemoveGroupingPolicies(filterPolicies)
				if err != nil {
					return err
				}

			}

			if serviceType == "service" {
				_, err = u.enforcer.RemoveGroupingPolicy(userStr, roleStr, domain)
				if err != nil {
					return err
				}
			}

			user := &model.User{Model: gorm.Model{ID: userID}}
			err = tx.Model(user).Association("Roles").Delete(&model.Role{Model: gorm.Model{ID: roleID}})
			if err != nil {
				return err
			}

			err = tx.Model(user).Association("Services").Delete(&model.Service{Model: gorm.Model{ID: serviceID}})

			if err != nil {
				return err
			}

			return nil
		})
	return
}

func (u *User) BindPermission(userID, roleID, serviceID uint) (err error) {
	db := u.core.GetDao().Storage().DB()
	userIDStr := userID2CasbinKey(userID)
	roleIDStr := roleID2CasbinKey(roleID)

	var domain string
	var serviceType string

	err = db.Transaction(
		func(tx *gorm.DB) error {
			service := &model.Service{}

			err = db.First(service, serviceID).Error
			if err != nil {
				return err
			}

			if service.ParentId != 0 {
				serviceType = "service"
				parentService := &model.Service{}
				err = db.First(parentService, service.ParentId).Error
				if err != nil {
					return err
				}

				domain = path.Join(strconv.Itoa(int(parentService.ID)), strconv.Itoa(int(service.ID)))
			} else {
				serviceType = "project"
				domain = filepath.Join(strconv.Itoa(int(service.ID)), "*")
			}

			if serviceType == "project" {
				filterPolicies := make([][]string, 0, 0)
				policies := u.enforcer.GetFilteredGroupingPolicy(0, userIDStr, roleIDStr)
				for _, v := range policies {
					has, err := filepath.Match(domain, v[2])
					if err != nil {
						return err
					}

					if has {
						filterPolicies = append(filterPolicies, v)
					}
				}

				_, err = u.enforcer.RemoveGroupingPolicies(filterPolicies)
				if err != nil {
					return err
				}

				_, err = u.enforcer.AddGroupingPolicy(userIDStr, roleIDStr, domain)
				if err != nil {
					return err
				}
			}

			if serviceType == "service" {
				_, err = u.enforcer.AddGroupingPolicy(userIDStr, roleIDStr, domain)
				if err != nil {
					return err
				}
			}
			user := &model.User{Model: gorm.Model{ID: userID}}
			err = tx.Model(user).Association("Roles").Append(&model.Role{Model: gorm.Model{ID: roleID}})
			if err != nil {
				return err
			}

			err = tx.Model(user).Association("Services").Append(&model.Service{Model: gorm.Model{ID: serviceID}})

			if err != nil {
				return err
			}

			return nil
		})
	return
}
