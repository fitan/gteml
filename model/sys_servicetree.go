/**
 * @Time : 3/4/21 6:05 PM
 * @Author : solacowa@gmail.com
 * @File : sys_namespace
 * @Software: GoLand
 */

package model

type TblServicetree struct {
	ID    int    `gorm:"column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Aname string `gorm:"column:aname" json:"aname"`
	Pnode int    `gorm:"column:pnode" json:"pnode"`
	// 1 namespace 2 service
	Type   int    `gorm:"column:type" json:"type"`
	Key    string `gorm:"column:key" json:"key"`
	Origin string `gorm:"column:origin" json:"origin"`
}

func (m *TblServicetree) TableName() string {
	return "tbl_servicetree"
}
