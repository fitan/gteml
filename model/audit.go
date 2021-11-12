/**
 * @Time : 2021/9/8 2:38 PM
 * @Author : solacowa@gmail.com
 * @File : audit
 * @Software: GoLand
 */

package model

import "time"

// Audit is the operation recoder
type Audit struct {
	ID                int       `json:"id,omitempty" gorm:"primary_key"`
	UserName          string    `json:"userName,omitempty"`
	Namespace         string    `json:"namespace,omitempty"`
	ClusterID         string    `json:"ClusterID,omitempty"`
	ResourceType      string    `json:"resourceType,omitempty"`
	ResourceReference string    `json:"resourceReference,omitempty"`
	Action            string    `json:"action,omitempty"`
	Operation         string    `json:"operation,omitempty"`
	StartTime         time.Time `json:"-" gorm:"-"`
	Duration          int64     `json:"duration,omitempty"`
	Status            int       `json:"status,omitempty"` //1 success 2 faild
	ExtData           []byte    `json:"extData,omitempty" gorm:"size:40960"`

	CreatedAt time.Time `json:"create_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
}

// TableName set table
func (*Audit) TableName() string {
	return "audit"
}
