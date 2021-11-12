/**
 * @Time : 8/5/21 3:07 PM
 * @Author : solacowa@gmail.com
 * @File : access
 * @Software: GoLand
 */

package model

import "time"

// 授权表
type AuthAccess struct {
	Id         int64      `gorm:"column:id;primary_key" json:"id"`
	Alias      string     `gorm:"column:alias;notnull;comment:'别名'" json:"alias"`
	AccessKey  string     `gorm:"column:access_key;notnull;unique;size:32;comment:'AccessKey'" json:"access_key"`
	SecretKey  string     `gorm:"column:secret_key;notnull;size:64;comment:'SecretKey'" json:"secret_key"`
	Namespace  string     `gorm:"column:namespace;notnull;index;size:24;comment:'项目空间'" json:"namespace"`
	ExpireTime *time.Time `gorm:"column:expire_time;null;comment:'过期时间'" json:"expireTime"` // 过期时间
	Enable     bool       `gorm:"column:enable;null;default:true;comment:'是否可用'" json:"enable"`
	Remark     string     `gorm:"column:remark;null;comment:'备注'" json:"remark"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"` // 更新时间
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`                   // 删除时间

	AuthServices  []AuthService  `gorm:"many2many:auth_access_service;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:auth_service_id;jointable_foreignkey:auth_access_id;" json:"authServices"`
	AuthResources []AuthResource `gorm:"many2many:auth_access_resource;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:auth_resource_id;jointable_foreignkey:auth_access_id;" json:"authResources"`
}

// TableName set table
func (*AuthAccess) TableName() string {
	return "auth_access"
}
