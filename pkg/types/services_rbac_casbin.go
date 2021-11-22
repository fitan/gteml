package types

type RBAC interface {
	GetRole() [][]string
	CreateRole(role string, obj string, method string) (err error)
	UpdateRole(oldRole, oldObj, oldMethod, newRole, newObj, newMethod string) (err error)
	DeleteRole(role string, obj string, method string) (err error)
	GetAuthorization() [][]string
	GetAuthorizationByUser(user string) [][]string
	AddAuthorization(user, role, domain string) (err error)
	DeleteAuthorization(user, role, domain string) (err error)
}
