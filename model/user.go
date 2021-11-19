/**
 * @Time : 3/4/21 5:26 PM
 * @Author : solacowa@gmail.com
 * @File : sys_user
 * @Software: GoLand
 */

package model

type User struct {
	//gorm.Model
	Id   int64
	Name string
	//唯一
	Email    string `gorm:"<-:create"`
	PassWord string
	Token    string
	Enable   bool
}

// TableName set table
func (u *User) TableName() string {
	return "user"
}
