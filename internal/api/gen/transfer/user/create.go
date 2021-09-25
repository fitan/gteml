package user

import (
	"github.com/fitan/gteml/internal/api/service/user"
	"github.com/fitan/gteml/pkg/core"
	"net/http"
)

type CreateTransfer struct {
	val user.CreateIn
}

func (t *CreateTransfer) Method() string {
	return http.MethodGet
}

func (t *CreateTransfer) Url() string {
	return "/user"
}

func (t *CreateTransfer) Binder() core.GinXBinder {
	return new(CreateBinder)
}

type CreateBinder struct {
	val user.CreateIn
}

func (b *CreateBinder) BindVal(core *core.Context) (interface{}, error) {
	err := core.GinX.ShouldBindUri(&b.val.Uri)
	b.val.Body.Hello = "anbowei"
	return b.val, err
}

func (b *CreateBinder) BindFn(core *core.Context) (interface{}, error) {
	return user.Create(core, &b.val)
}
