package types

type Version interface {
	AddVersion()
	Version() int
}
