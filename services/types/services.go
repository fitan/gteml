package types

type Serviceser interface {
	User() Userer
	RABC() RBAC
	Audit() Audit
	K8s() K8s
}
