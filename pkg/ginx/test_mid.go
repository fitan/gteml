package ginx

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/pkg/errors"
)

type TestMid struct {
}

func (t *TestMid) BindValBefor(core *types.Core) bool {
	core.GinX.SetBindErr(errors.New("TestMid: BindValBefor"))
	return true
}

func (t *TestMid) BindValAfter(core *types.Core) bool {
	core.GinX.SetBindErr(errors.New("TestMid: BindValAfter"))
	return true
}

func (t *TestMid) BindFnAfter(core *types.Core) bool {
	core.GinX.SetBindErr(errors.New("TestMid: BindFnAfter"))
	return true
}

func (t *TestMid) Forever(core *types.Core) {
	tempErr := core.GinX.BindErr()
	for {
		if tempErr == nil {
			return
		}
		tempErr = errors.Unwrap(tempErr)
	}
}
