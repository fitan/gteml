/**
 * @Time : 3/4/21 5:26 PM
 * @Author : solacowa@gmail.com
 * @File : sys_user
 * @Software: GoLand
 */

package model

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	gorm.Model
	Name string
	//唯一
	Email    string
	PassWord string
	Token    string
	Enable   bool

	Roles    []Role    `gorm:"many2many:user_roles"`
	Services []Service `gorm:"many2many:user_services"`
}

type Method interface {
	// Where("id=@id")
	GetByID(id uint) (gen.T, error)
}
