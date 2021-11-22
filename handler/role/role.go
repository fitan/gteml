package role

import (
	"github.com/fitan/magic/pkg/types"
	"strconv"
)

type AddPermissionsIn struct {
	Uri struct {
		UserId int64 `json:"user_id"`
	} `json:"uri"`
	Body Domain `json:"body"`
}

func AddPermissions(core *types.Core, in *AddPermissionsIn) (string, error) {
}

type GetRolesForUserInDomainIn struct {
	Url struct {
		UserId int64 `uri:"user_id" json:"user_id"`
	}
	Query Domain `json:"query"`
}

func (g *GetRolesForUserInDomainIn) DomainConv() string {
	return g.Query.DomainConv()
}

func GetRolesForUserInDomain(core *types.Core, in *GetRolesForUserInDomainIn) {
	core.GetServices().RABC().GetRolesForUserInDomain(strconv.FormatInt(in.Url.UserId, 10), in.DomainConv())
}
