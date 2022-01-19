package types

import "github.com/fitan/magic/dao/dal/model"

type Role interface {
	//
	//  UnBindPermission
	//  @Description: 接触绑定
	//  @param roleID
	//  @param permissionID
	//  @return err
	//
	UnBindPermission(roleID uint, permissionID uint) (err error)
	BindPermission(roleID uint, permissionID uint) (err error)
	Get() (res []*model.Role, err error)
	GetById(id uint) (res *model.Role, err error)
	Create(role *model.Role) error
	DeleteById(id uint) (err error)
}
