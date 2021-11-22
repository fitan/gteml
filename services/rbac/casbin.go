package rbac

import (
	"github.com/casbin/casbin/v2"
	"github.com/fitan/magic/pkg/types"
	"go.uber.org/zap"
)

func NewRBAC(enforcer *casbin.Enforcer, core types.ServiceCore) *RBAC {
	return &RBAC{
		core:     core,
		enforcer: enforcer,
	}

}

type RBAC struct {
	core     types.ServiceCore
	enforcer *casbin.Enforcer
}

func (r *RBAC) GetRole() [][]string {
	return r.enforcer.GetPolicy()
}

func (r *RBAC) CreateRole(role string, obj string, method string) (err error) {
	log := r.core.GetCoreLog().ApmLog("services.rbac.createRole")
	defer func() {
		if err != nil {
			log.Error(
				err.Error(), zap.Any("methodIn", map[string]interface{}{"role": role, "obj": obj, "method": method}),
			)
		}

		log.Sync()
	}()

	_, err = r.enforcer.AddPolicy(role, obj, method)
	return
}

func (r *RBAC) UpdateRole(oldRole, oldObj, oldMethod, newRole, newObj, newMethod string) (err error) {
	log := r.core.GetCoreLog().ApmLog("services.rbac.updateRole")
	defer func() {
		if err != nil {
			log.Error(
				err.Error(), zap.Any(
					"methodIn", map[string]interface{}{
						"oldRole": oldRole, "oldObj": oldObj, "oldMethod": oldMethod, "newRole": newRole, "newObj": newObj, "newMethod": newMethod,
					},
				),
			)
		}

		log.Sync()
	}()

	_, err = r.enforcer.UpdatePolicy([]string{oldRole, oldObj, oldMethod}, []string{newRole, newObj, newMethod})
	return
}

func (r *RBAC) DeleteRole(role string, obj string, method string) (err error) {
	log := r.core.GetCoreLog().ApmLog("services.rbac.deleteRole")
	defer func() {
		if err != nil {
			log.Error(
				err.Error(), zap.Any("methodIn", map[string]interface{}{"role": role, "obj": obj, "method": method}),
			)
		}

		log.Sync()
	}()

	_, err = r.enforcer.RemovePolicy(role, obj, method)
	return
}

func (r *RBAC) GetAuthorization() [][]string {
	return r.enforcer.GetGroupingPolicy()
}

func (r *RBAC) GetAuthorizationByUser(user string) [][]string {
	return r.enforcer.GetFilteredGroupingPolicy(1, user)
}

func (r *RBAC) AddAuthorization(user, role, domain string) (err error) {
	log := r.core.GetCoreLog().ApmLog("services.rbac.addAuthorization")
	defer func() {
		if err != nil {
			log.Error(
				err.Error(), zap.Any("methodIn", map[string]interface{}{"user": user, "role": role, "domain": domain}),
			)
		}

		log.Sync()
	}()
	_, err = r.enforcer.AddGroupingPolicy(user, role, domain)
	return
}

func (r *RBAC) DeleteAuthorization(user, role, domain string) (err error) {
	log := r.core.GetCoreLog().ApmLog("services.rbac.deleteAuthorization")
	defer func() {
		if err != nil {
			log.Error(
				err.Error(), zap.Any("methodIn", map[string]interface{}{"user": user, "role": role, "domain": domain}),
			)
		}

		log.Sync()
	}()

	_, err = r.enforcer.RemoveGroupingPolicy(user, role, domain)
	return
}
