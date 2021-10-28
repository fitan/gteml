package user

import (
	"github.com/fitan/magic/internal/api/service/user"
	"github.com/fitan/magic/pkg/types"
	"net/http"
)

type CreateTransfer struct {
}

func (t *CreateTransfer) Method() string {
	return http.MethodGet
}

func (t *CreateTransfer) Url() string {
	return "/user"
}

func (t *CreateTransfer) Binder() types.GinXBinder {
	return new(CreateBinder)
}

type CreateBinder struct {
	val user.CreateIn
}

func (b *CreateBinder) BindVal(core *types.Context) (interface{}, error) {
	err := core.GinX.GinCtx().ShouldBindUri(&b.val.Uri)
	b.val.Body.Hello = "anbowei"
	return b.val, err
}

func (b *CreateBinder) BindFn(core *types.Context) (interface{}, error) {
	return user.Create(core, &b.val)
}
