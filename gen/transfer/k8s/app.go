package k8s

import (
	"net/http"

	"github.com/fitan/magic/handler/k8s"
	"github.com/fitan/magic/pkg/types"
)

type GetAppTransfer struct {
}

func (t *GetAppTransfer) Method() string {
	return http.MethodGet
}

func (t *GetAppTransfer) Url() string {
	return "/k8s/:namespace/app/:name"
}

func (t *GetAppTransfer) Binder() types.GinXBinder {
	return new(GetAppBinder)
}

type GetAppBinder struct {
	val *k8s.GetAppIn
}

func (b *GetAppBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &k8s.GetAppIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Success 200 {object} ginx.GinXResult{data=v1beta1.Application}
// @Description 获取app
// @Router /k8s/:namespace/app/:name [get]
func (b *GetAppBinder) BindFn(core *types.Core) (interface{}, error) {
	return k8s.GetApp(core, b.val)
}
