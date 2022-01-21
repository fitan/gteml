package storage

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/dao/dal/model"
	query2 "github.com/fitan/magic/dao/dal/query"
	types2 "github.com/fitan/magic/dao/types"
	"github.com/fitan/magic/pkg/dbquery"
	"github.com/fitan/magic/pkg/types"
	"github.com/pkg/errors"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"path"
	"path/filepath"
	"strconv"
)

func NewUser(query *dbquery.WrapQuery, daoCore types.ServiceCore, enforcer *casbin.Enforcer) types2.User {
	return &User{query: query, core: daoCore, enforcer: enforcer}
}

type User struct {
	query    *dbquery.WrapQuery
	core     types.ServiceCore
	enforcer *casbin.Enforcer
}

func (u *User) CheckUserPermission(userID uint, serviceID uint, path, method string) (err error) {
	service, err := u.query.WrapQuery().Service.Where(u.query.Service.ID.Eq(serviceID)).First()
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
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("permission denied")
	}
	return nil
}

func (u *User) CheckPassword(userName string, password string) (res *model.User, err error) {
	//ctx := u.core.GetTrace().ApmSpanCtx("sql checkpssword")
	//return u.query.RawQ().WithContext(ctx).User.Where(u.query.User.Email.Eq(userName)).Where(u.query.User.PassWord.Eq(password)).First()
	return u.query.WrapQuery().User.Where(u.query.User.Email.Eq(userName), u.query.User.PassWord.Eq(password)).First()
}

func (u *User) ById(id uint, preload ...field.RelationField) (*model.User, error) {
	do := u.query.WrapQuery().User.Where(u.query.User.ID.Eq(id))
	if len(preload) != 0 {
		for _, v := range preload {
			do = do.Preload(v)
		}
	}
	return do.First()
}

func (u *User) Create(user *model.User) error {
	return u.query.WrapQuery().User.Create(user)
}

func (u *User) Update(id uint, user *model.User) error {
	_, err := u.query.WrapQuery().User.Where(u.query.User.ID.Eq(id)).Updates(user)
	return err
}

func (u *User) UnBindPermission(userID, roleID, serviceID uint) (err error) {
	var domain string
	var serviceType string

	err = u.query.Transaction(
		func(tx *query2.Query) error {
			service, err := tx.Service.Where(tx.Service.ID.Eq(serviceID)).First()
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
			err = tx.User.Roles.Model(user).Delete(&model.Role{Model: gorm.Model{ID: roleID}})
			//err = tx.Model(user).Association("Roles").Delete(&model.Role{Model: gorm.Model{ID: roleID}})
			if err != nil {
				return err
			}

			//err = tx.Model(user).Association("Services").Delete(&model.Service{Model: gorm.Model{ID: serviceID}})
			err = tx.User.Services.Model(user).Delete(&model.Service{Model: gorm.Model{ID: serviceID}})

			if err != nil {
				return err
			}

			return nil

		})

	return
}

func (u *User) BindPermission(userID, roleID, serviceID uint) (err error) {
	userIDStr := userID2CasbinKey(userID)
	roleIDStr := roleID2CasbinKey(roleID)

	var domain string
	var serviceType string

	err = u.query.Transaction(
		func(tx *query2.Query) error {

			service, err := tx.Service.Where(tx.Service.ID.Eq(serviceID)).First()
			if err != nil {
				return err
			}

			if service.ParentId != 0 {
				serviceType = "service"
				parentService, err := tx.Service.Where(tx.Service.ID.Eq(service.ParentId)).First()
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
			err = tx.User.Roles.Model(user).Append(&model.Role{Model: gorm.Model{ID: roleID}})
			if err != nil {
				return err
			}

			err = tx.User.Services.Model(user).Append(&model.Service{Model: gorm.Model{ID: serviceID}})

			if err != nil {
				return err
			}

			return nil
		})

	return
}
