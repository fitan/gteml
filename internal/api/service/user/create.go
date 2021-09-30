package user

import (
	"errors"
	"github.com/fitan/gteml/pkg/core"
)

type CreateIn struct {
	Body struct {
		Hello string `json:"hello"`
	} `json:"body"`
	Uri    struct{}
	Header struct{}
}

// @Router post /user
func Create(c *core.Context, in *CreateIn) (string, error) {
	c.Log.Info("这是 create的开始")

	c.Apis.Baidu.GetRoot()

	res, ok := c.GinX.GetQuery("status")
	if !ok {
		return "", errors.New("not find query status")
	}
	return res, nil
}
