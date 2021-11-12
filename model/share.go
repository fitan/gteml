/**
 * @Time : 7/5/21 5:19 PM
 * @Author : solacowa@gmail.com
 * @File : share
 * @Software: GoLand
 */

package model

import "time"

// 系统组，与namespace类似，预留还没想好怎么用
type Share struct {
	Id         int64      `gorm:"column:id;primary_key" json:"id"`
	Code       string     `gorm:"column:code;notnull;unique;size:24;comment:'分享唯一标识'" json:"code"`     // 分享标识
	Link       string     `gorm:"column:link;notnull;size:1024;comment:'目标路径'" json:"link"`            // 目标路径
	Bucket     string     `gorm:"column:bucket;notnull;index;size:128;comment:'Bucket'" json:"bucket"` // Bucket
	ExpireTime *time.Time `gorm:"column:expireTime;null;comment:'过期时间'" json:"expireTime"`             // 过期时间
	Public     bool       `gorm:"column:public;null;comment:'是否公开'" json:"public"`                     // 是否公开
	Password   string     `gorm:"column:password;null;comment:'查询密码'" json:"password"`                 // 查询密码
	IsDir      bool       `gorm:"column:isDir;null;default:false;comment:'是否是目录'" json:"isDir"`        // 是否是目录
	Count      int64      `gorm:"column:count;null;comment:'引用计数'" json:"count"`                       // 引用计数
	Sharer     string     `gorm:"column:sharer;notnull;index;comment:'分享人'" json:"sharer"`             // 分享人
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"`               // 创建时间
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`               // 更新时间
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`                                 // 删除时间
}

// TableName set table
func (*Share) TableName() string {
	return "shares"
}
