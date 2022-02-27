/**
 * @Time : 3/4/21 5:26 PM
 * @Author : solacowa@gmail.com
 * @File : sys_user
 * @Software: GoLand
 */

package model

import (
	"gorm.io/gen"
)

type User struct {
	Model
	Name string `json:"name"`
	//唯一
	Email    string `json:"email"`
	PassWord string `json:"passWord"`
	Token    string `json:"token"`
	Enable   bool   `json:"enable"`

	Roles    []Role    `gorm:"many2many:user_roles" json:"roles"`
	Services []Service `gorm:"many2many:user_services" json:"services"`
}

type ApiUser struct {
	Model
	Name string
	//唯一
	Email  string
	Token  string
	Enable bool
}

type QueryUser struct {
	Name string `json:"name" gen:"="`
	Or   struct {
		Email string `json:"email" gen:"="`
	}
	Paging struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}
	Roles struct {
		Id int32 `json:"id" gen:"="`
	} `json:"roles" gen:"relation"`
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
