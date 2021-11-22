/**
 * @Time : 3/4/21 6:05 PM
 * @Author : solacowa@gmail.com
 * @File : sys_permission
 * @Software: GoLand
 */

package model

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ParentId    int64  `gorm:"column:parent_id;null;default:0;comment:'上级ID'" json:"parent_id"`
	Icon        string `gorm:"column:icon;null;comment:'Icon'" json:"icon"`
	Menu        bool   `gorm:"column:menu;null;default:false;comment:'是否是菜单'" json:"menu"`
	Method      string `gorm:"column:method;null;default:'GET';comment:'Method'" json:"method"`
	Alias       string `gorm:"column:alias;notnull;comment:'名称'" json:"alias"`
	Name        string `gorm:"column:name;unique;size:32;notnull;comment:'标识'" json:"name"`
	Path        string `gorm:"column:path;unique;size:64;notnull;comment:'路径'" json:"path"`
	Description string `gorm:"column:description;null;comment:'描述'" json:"description"`
}
