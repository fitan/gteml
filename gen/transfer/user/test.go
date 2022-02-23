package user

import (
	"net/http"

	"github.com/fitan/magic/handler/user"
	"github.com/fitan/magic/pkg/ginbind"
	"github.com/fitan/magic/pkg/types"
	"github.com/gin-gonic/gin/binding"
)

type RestUsersTransfer struct {
}

func (t *RestUsersTransfer) Method() string {
	return http.MethodGet
}

func (t *RestUsersTransfer) Url() string {
	return "/restful/users"
}

func (t *RestUsersTransfer) FuncName() string {
	return "user.RestUsers"
}

func (t *RestUsersTransfer) Binder() types.GinXBinder {
	return new(RestUsersBinder)
}

type RestUsersBinder struct {
	val *user.RestUsersIn
}

func (b *RestUsersBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.RestUsersIn{}
	b.val = in

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Success 200 {object} ginx.XResult{data=[]model.User}
// @Description 获取Users
// @Router /restful/users [get]
func (b *RestUsersBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.RestUsers(core, b.val)
}

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
// @Success 200 {object} ginx.XResult{data=string}
// @Router /user [post]
func (b *CreateBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.Create(core, b.val)
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

func (t *SayHelloTransfer) FuncName() string {
	return "user.SayHello"
}

func (t *SayHelloTransfer) Binder() types.GinXBinder {
	return new(SayHelloBinder)
}

type SayHelloBinder struct {
	val *user.SayHelloIn
}

func (b *SayHelloBinder) BindVal(core *types.Core) (res interface{}, err error) {
	in := &user.SayHelloIn{}
	b.val = in

	err = core.GinX.GinCtx().ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	err = ginbind.BindCtxKey(core.GinX.GinCtx(), &in.CtxKey)
	if err != nil {
		return nil, err
	}

	err = binding.Validator.ValidateStruct(&in.CtxKey)
	if err != nil {
		return nil, err
	}

	return b.val, err
}

// @Accept  json
// @Produce  json
// @Param query query SwagSayHelloQuery false " "
// @Success 200 {object} ginx.XResult{data=string}
// @Router /say [get]
func (b *SayHelloBinder) BindFn(core *types.Core) (interface{}, error) {
	return user.SayHello(core, b.val)
}
