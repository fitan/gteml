/**
 * @Time : 3/4/21 5:26 PM
 * @Author : solacowa@gmail.com
 * @File : sys_user
 * @Software: GoLand
 */

package model

import "time"

type SysUser struct {
	Active         int       `gorm:"column:active"`
	APIToken       string    `gorm:"column:api_token"`
	AuthType       int       `gorm:"column:auth_type"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	Displayname    string    `gorm:"column:displayname"`
	Email          string    `gorm:"column:email"`
	Id             int64     `gorm:"column:id;primary_key"`
	LoginFrequency int       `gorm:"column:login_frequency"`
	Name           string    `gorm:"column:name"`
	OnlyOss        int       `gorm:"column:only_oss"`
	Password       string    `gorm:"column:password"`
	Phone          string    `gorm:"column:phone"`
	Role           int       `gorm:"column:role"`
	UpdatedAt      time.Time `gorm:"column:updated_at"` // 删除时间

	SysRoles       []SysRole        `gorm:"many2many:sys_user_roles;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:role_id;jointable_foreignkey:sys_user_id;" json:"sys_roles"`
	SysServicetree []TblServicetree `gorm:"many2many:sys_user_servicetree;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:servicetree_id;jointable_foreignkey:sys_user_id;" json:"sys_servicetree"`
}

// TableName set table
func (*SysUser) TableName() string {
	return "user"
}
