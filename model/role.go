/**
 * @Time : 3/4/21 6:04 PM
 * @Author : solacowa@gmail.com
 * @File : role
 * @Software: GoLand
 */

package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string
	OnlyKey     string
	Enabled     bool
	Description string
	Level       int
	Permissions []Permission `gorm:"many2many:role_permissions"`
}
