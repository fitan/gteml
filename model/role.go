/**
 * @Time : 3/4/21 6:04 PM
 * @Author : solacowa@gmail.com
 * @File : role
 * @Software: GoLand
 */

package model

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	Name        string `gorm:"column:name;notnull;size:24;unique;comment:'标识'" json:"name"`
	OnlyKey     string
	Enabled     bool         `gorm:"column:enabled;null;default:true;comment:'是否可用'" json:"enabled"`
	Description string       `gorm:"column:description;null;comment:'描述'" json:"description"`
	Level       int          `gorm:"column:level;null;comment:'等级'" json:"level"`
	Permissions []Permission `gorm:"many2many:sys_role_permissions;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:permission_id;jointable_foreignkey:role_id;"`
}

// TableName sets the insert table name for this struct type
func (r *SysRole) TableName() string {
	return "sys_role"
}
