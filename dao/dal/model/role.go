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
	Name        string       `json:"name"`
	OnlyKey     string       `json:"onlyKey"`
	Enabled     bool         `json:"enabled"`
	Description string       `json:"description"`
	Level       int          `json:"level"`
	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions"`
}
