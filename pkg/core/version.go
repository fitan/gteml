package core

import (
	"github.com/fitan/gteml/pkg/types"
	"sync"
)

type Version struct {
	m sync.Mutex
	v int
}

func NewVersion() types.Version {
	return &Version{
		m: sync.Mutex{},
		v: 0,
	}
}

func (v *Version) AddVersion() {
	v.m.Lock()
	defer v.m.Unlock()
	v.v = v.v + 1
}

func (v *Version) Version() int {
	v.m.Lock()
	defer v.m.Unlock()
	return v.v
}

type VersionReg struct {
	version types.Version
}

func (v *VersionReg) Reload(c *Context) {
	panic("implement me")
}

func (v *VersionReg) With(o ...Option) Register {
	panic("implement me")
}

func (v *VersionReg) Set(c *Context) {
	if v.version == nil {
		v.version = NewVersion()
	}
	c.Version = v.version
}

func (v *VersionReg) Unset(c *Context) {
}
