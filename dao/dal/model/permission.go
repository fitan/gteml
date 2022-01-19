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
	ParentId    uint
	Icon        string
	Menu        bool
	Method      string
	Alias       string
	Name        string
	Path        string
	Description string
}
