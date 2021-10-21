package types

type Version interface {
	AddVersion()
	Version() int
	SetVersion(c *Context)
	CompareVersion(c *Context) bool
}
