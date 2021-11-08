package user

import (
	"github.com/fitan/magic/pkg/types"
	user2 "github.com/fitan/magic/service/user"
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
	val user2.CreateIn
}

func (b *CreateBinder) BindVal(core *types.Core) (interface{}, error) {
	err := core.GinX.GinCtx().ShouldBindUri(&b.val.Uri)
	b.val.Body.Hello = "anbowei"
	return b.val, err
}

func (b *CreateBinder) BindFn(core *types.Core) (interface{}, error) {
	return user2.Create(core, &b.val)
}

type SayHelloTransfer struct {
}

func (s *SayHelloTransfer) Method() string {
	return http.MethodGet
}

func (s *SayHelloTransfer) Url() string {
	return "/say"
}

func (s *SayHelloTransfer) Binder() types.GinXBinder {
	return &SayHelloBinderBinder{}
}

type SayHelloBinderBinder struct {
	val user2.SayHelloIn
}

func (s *SayHelloBinderBinder) BindVal(c *types.Core) (interface{}, error) {
	err := c.GinX.GinCtx().BindQuery(&s.val.Query)
	return s.val, err
}

func (s *SayHelloBinderBinder) BindFn(c *types.Core) (interface{}, error) {
	return user2.SayHello(c, &s.val)
}
