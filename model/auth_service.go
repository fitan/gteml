/**
 * @Time : 8/5/21 3:12 PM
 * @Author : solacowa@gmail.com
 * @File : auth_access_service
 * @Software: GoLand
 */

package model

import "time"

// 授权的哪些服务
type AuthService struct {
	Id        int64      `gorm:"column:id;primary_key" json:"id"`
	Name      string     `gorm:"column:name;notnull;index;size:24;comment:'服务名称'" json:"name"`
	Enable    bool       `gorm:"column:enable;null;default:true;comment:'是否可用'" json:"enable"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"` // 更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`                   // 删除时间
}

// TableName set table
func (*AuthService) TableName() string {
	return "auth_service"
}
