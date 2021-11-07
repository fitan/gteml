package types

type Version interface {
	AddVersion()
	Version() int
	SetVersion(c *Core)
	CompareVersion(c *Core) bool
}
