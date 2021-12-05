package ginbind

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"reflect"
)

func BindCtxKey(ctx *gin.Context, i interface{}) error {
	iV := reflect.ValueOf(i)
	if iV.Type().Kind() != reflect.Ptr {
		return errors.New("binding value not ptr")
	}

	iV = iV.Elem()

	for i := 0; i < iV.Type().NumField(); i++ {
		tf := iV.Type().Field(i)

		if tf.Anonymous {
			continue
		}

		tagV, ok := tf.Tag.Lookup("ctxkey")
		if !ok {
			continue
		}

		value, ok := ctx.Get(tagV)

		if ok {
			valueV := reflect.ValueOf(value)
			if iV.Field(i).Type() != valueV.Type() {
				return fmt.Errorf("ctxkey field %v type not equal to %v", iV.Field(i).Type().Name(), valueV.Type().String())
			}
			iV.Field(i).Set(valueV)
		}
	}
	return nil
}
