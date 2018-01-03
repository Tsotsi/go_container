package main

import (
	"reflect"
	"go-container"
	"fmt"
)

type Abc struct {
	Cc int
}

func main() {
	c := go_container.NewContainer()
	ab := reflect.TypeOf(c)
	fmt.Println(ab.Name())
	fmt.Print()
	c.Set("abc", func(cc *go_container.Container) interface{} {
		a := new(Abc)
		a.Cc = 1988
		return a
	})
	if f, ok := c.Get("abc"); ok {
		tt := reflect.TypeOf(f)
		fmt.Println(tt.Key().Name())
	}
}
