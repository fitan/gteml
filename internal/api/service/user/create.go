package user

import "github.com/fitan/gteml/pkg/core"

type CreateTransfer struct {
	val CreateIn
}

func (c *CreateTransfer) Method() string {
	return "Post"
}

func (c *CreateTransfer) Url() string {
	return "/user"
}

func (c *CreateTransfer) BindVal(core *core.Context) (interface{}, error) {
	err := core.GinX.ShouldBindUri(&c.val.Uri)
	return c.val, err
}

func (c *CreateTransfer) BindFn(core *core.Context) (interface{}, error) {
	return Create(core, &c.val)
}

type CreateIn struct {
	Body   struct{}
	Uri    struct{}
	Header struct{}
}

func Create(c *core.Context, in *CreateIn) (*CreateIn, error) {
	return in, nil
}
