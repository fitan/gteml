package core

import (
	"github.com/fitan/magic/pkg/types"
)

type Version struct {
	v int
}

func NewVersion() types.Version {
	return &Version{
		v: 0,
	}
}

func (v *Version) AddVersion() {
	v.v = v.v + 1
}

func (v *Version) Version() int {
	return v.v
}

func (v *Version) SetVersion(c *types.Core) {
	c.LocalVersion = v.Version()
}

func (v *Version) CompareVersion(c *types.Core) bool {
	return v.Version() == c.LocalVersion
}

type VersionRegister struct {
	version types.Version
}

func (v *VersionRegister) Reload(c *types.Core) {
}

func (v *VersionRegister) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (v *VersionRegister) Set(c *types.Core) {
	if v.version == nil {
		v.version = NewVersion()
	}
	c.Version = v.version
}

func (v *VersionRegister) Unset(c *types.Core) {
}
