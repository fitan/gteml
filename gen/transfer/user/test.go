package user

import (
	"net/http"

	"github.com/fitan/magic/handler/user"
	"github.com/fitan/magic/pkg/types"
)

type SwagCreateBody struct {
	Hello string `json:"hello"`
}

type CreateTransfer struct {
}

func (t *CreateTransfer) Method() string {
	return http.MethodPost
}

func (t *CreateTransfer) Url() string {
	return "/user"
}

func (t *CreateTransfer) FuncName() string {
	return "user.Create"
}

func (t *CreateTransfer) Binder() types.GinXBinder {
	return new(CreateBinder)
}

type CreateBinder struct {
	val *user.CreateIn
}

func (b *CreateBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.CreateIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param body body SwagCreateBody true " "
// @Success 200 {object} ginx.GinXResult{data=string}
// @Router /user [post]
func (b *CreateBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.Create(core, b.val)
}
