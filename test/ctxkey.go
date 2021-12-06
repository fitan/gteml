package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name     string
	Password string
}

type C struct {
	CtxKey anonymous
}

type anonymous struct {
	ID   int  `ctxkey:"id"`
	User User `ctxkey:"user"`
}

func CtxKey() {
	keyM := make(map[string]interface{}, 0)
	keyM["id"] = 1
	keyM["user"] = User{
		Name:     "boweian",
		Password: "123456",
	}

	c := &C{}

	iV := reflect.ValueOf(&c.CtxKey).Elem()
	for i := 0; i < iV.Type().NumField(); i++ {
		tf := iV.Type().Field(i)

		if tf.Anonymous {
			continue
		}

		tagV, ok := tf.Tag.Lookup("ctxkey")
		if !ok {
			continue
		}

		key, ok := keyM[tagV]
		if ok && iV.Field(i).Type() == reflect.TypeOf(key) {
			iV.Field(i).Set(reflect.ValueOf(key))
		} else {
			fmt.Println(iV.Field(i).Type().Name()+" type not equal to: ", reflect.TypeOf(key))
		}
	}
	fmt.Println(c)
}

func main() {
	CtxKey()
}

type Say struct {
	Say2 Say2
}

func (s *Say) Say1() {

}

type Say2 struct {
}

func (s *Say2) Say2() {

}

type Sayer interface {
	Say1()
	Say2()
}
