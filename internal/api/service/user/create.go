package user

import (
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
func Create(c *core.Context, in *CreateIn) (*CreateIn, error) {
	return in, nil
}
