package user

import (
	"errors"
	"github.com/fitan/gteml/pkg/types"
)

type CreateIn struct {
	Body struct {
		Hello string `json:"hello"`
	} `json:"body"`
	Uri    struct{}
	Header struct{}
}

// @Router post /user
func Create(c *types.Context, in *CreateIn) (string, error) {
	c.Log.Info("这是 create的开始")

	data, _ := c.Apis.Baidu().GetRoot()

	_, ok := c.GinX.GinCtx().GetQuery("status")
	if !ok {
		return "", errors.New("not find query status")
	}
	return data.String(), nil
}
