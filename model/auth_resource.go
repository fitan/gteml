/**
 * @Time : 8/5/21 3:14 PM
 * @Author : solacowa@gmail.com
 * @File : auth_method
 * @Software: GoLand
 */

package model

import "time"

// AuthResource 可操作的资源
type AuthResource struct {
	Id        int64      `gorm:"column:id;primary_key" json:"id"`
	Name      string     `gorm:"column:name;notnull;unique;size:24;comment:'名称'" json:"name"`
	Alias     string     `gorm:"column:alias;notnull;size:64;comment:'中文名称'" json:"alias"`
	Path      string     `gorm:"column:path;notnull;unique;size:128;comment:'API路径'" json:"path"`
	Method    string     `gorm:"column:method;null;default:'GET';size:12;comment:'Method方法'" json:"method"`
	Enable    bool       `gorm:"column:enable;null;default:true;comment:'是否可用'" json:"enable"`
	Style     string     `gorm:"column:style;null;comment:'资源类型nginx资源'" json:"style"` // 资源类型
	Docs      string     `gorm:"column:docs;null;type:text;comment:'说明文档'" json:"docs"`
	Remark    string     `gorm:"column:remark;null;comment:'备注'" json:"remark"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"` // 更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`                   // 删除时间
}

// TableName set table
func (*AuthResource) TableName() string {
	return "auth_resource"
}
