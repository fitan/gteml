package role

import "strconv"

type Domain struct {
	ProjectID int64 `form:"project_id" json:"project_id"`
	ServiceID int64 `form:"service_id" json:"service_id"`
}

func (d *Domain) DomainConv() string {
	return strconv.FormatInt(d.ProjectID, 10) + "." + strconv.FormatInt(d.ServiceID, 10)
}
