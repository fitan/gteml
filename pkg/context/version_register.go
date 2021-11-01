package context

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

func (v *Version) SetVersion(c *types.Context) {
	c.LocalVersion = v.Version()
}

func (v *Version) CompareVersion(c *types.Context) bool {
	return v.Version() == c.LocalVersion
}

type VersionReg struct {
	version types.Version
}

func (v *VersionReg) Reload(c *types.Context) {
}

func (v *VersionReg) With(o ...types.Option) types.Register {
	panic("implement me")
}

func (v *VersionReg) Set(c *types.Context) {
	if v.version == nil {
		v.version = NewVersion()
	}
	c.Version = v.version
}

func (v *VersionReg) Unset(c *types.Context) {
}
