/**
 * @Time : 2021/11/3 10:44 AM
 * @Author : solacowa@gmail.com
 * @File : tbl_authority
 * @Software: GoLand
 */

package model

type TblAuthority struct {
	User            string `json:"user,omitempty"`
	NodeId          int    `json:"nodeid,omitempty"`
	NodeName        string `json:"nodename,omitempty"`
	NodeAname       string `json:"nodeaname,omitempty"`
	Key             string `json:"key,omitempty"`
	ParentnodeId    int    `json:"parentnodeid,omitempty"`
	ParentnodeName  string `json:"parentnodename,omitempty"`
	ParentnodeAname string `json:"parentnodeaname,omitempty"`
	RoleId          int    `json:"roleid,omitempty"`

	UserInfo SysUser `gorm:"foreignkey:User;references:Email;association_foreignkey:Email;OnDelete:SET NULL;"`
}

func (*TblAuthority) TableName() string {
	return "tbl_authority"
}
