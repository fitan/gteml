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

type ApiUser struct {
	gorm.Model
	Name string
	//唯一
	Email  string
	Token  string
	Enable bool
}

type UserMethod interface {
	// where("id=@id")
	GetByID(id uint) (gen.T, error)
	// where("email=@email and pass_word=@password")
	CheckAccount(email, password string) (gen.T, error)
	// update @@table {{set}} pass_word=@password {{end}} {{where}} id=@id {{end}}
	ModifyPassword(id int, password string) error
	// select * from @@table
	FindApi() ([]ApiUser, error)
}
