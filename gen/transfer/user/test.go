package user

import (
	"github.com/fitan/magic/handler/user"
	"github.com/fitan/magic/pkg/types"
	"net/http"
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

func (t *CreateTransfer) Binder() types.GinXBinder {
	return new(CreateBinder)
}

type CreateBinder struct {
	val user.CreateIn
}

func (b *CreateBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.CreateIn{}

	err = core.GinX.GinCtx().ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	err = core.GinX.GinCtx().ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	err = core.GinX.GinCtx().ShouldBindHeader(&in.Header)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param body body SwagCreateBody true " "
// @Success 200 {object} public.Result{data=[]ent.User}
// @Router /user [post]
func (b *CreateBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.Create(core, &b.val)
}

type SwagSayHelloQuery struct {
	Say string `json:"say" form:"say"`
}

type SayHelloTransfer struct {
}

func (t *SayHelloTransfer) Method() string {
	return http.MethodGet
}

func (t *SayHelloTransfer) Url() string {
	return "/say"
}

func (t *SayHelloTransfer) Binder() types.GinXBinder {
	return new(SayHelloBinder)
}

type SayHelloBinder struct {
	val user.SayHelloIn
}

func (b *SayHelloBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.SayHelloIn{}

	err = core.GinX.GinCtx().ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param query query SwagSayHelloQuery false " "
// @Success 200 {object} public.Result{data=string}
// @Router /say [get]
func (b *SayHelloBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.SayHello(core, &b.val)
}