package go_container

import (
	"testing"
	"fmt"
)

type Abc struct {
	Cc int
}

func NewAbc()  *Abc{
	c:=new(Abc)
	c.Cc = 110
	return c
}

const N = 10

func TestContainer(t *testing.T) {
	c := NewContainer()
	c.Set("abc", func(cc *Container) interface{} {
		return NewAbc()
	})
	if f, ok := c.Get("abc"); ok {
		if f.(*Abc).Cc == 110 {
			t.Log(f.(*Abc).Cc)
		} else {
			t.Error(f.(*Abc).Cc)
		}
		f.(*Abc).Cc += 9
	}

	if f, ok := c.Get("abc"); ok {
		if f.(*Abc).Cc == 119 {
			t.Log(f.(*Abc).Cc)
		} else {
			t.Error(f.(*Abc).Cc)
		}
	}

}


func TestContainer_Factory(t *testing.T) {
	c := NewContainer()
	c.Factory("abc", func(cc *Container) interface{} {
		return NewAbc()
	})
	if f, ok := c.Get("abc"); ok {
		if f.(*Abc).Cc == 110 {
			t.Log(f.(*Abc).Cc)
		} else {
			t.Error(f.(*Abc).Cc)
		}
		f.(*Abc).Cc += 9
	}

	if f, ok := c.Get("abc"); ok {
		if f.(*Abc).Cc == 110 {
			t.Log(f.(*Abc).Cc)
		} else {
			t.Error(f.(*Abc).Cc)
		}
	}

}

type exampleAbc struct {
	Cc int
}

func newExampleAbc()  *exampleAbc{
	c:=new(exampleAbc)
	c.Cc = 110
	return c
}

func ExampleContainer_Set()  {
	c := NewContainer()
	c.Set("abc", func(cc *Container) interface{} {
		return newExampleAbc()
	})
	if f, ok := c.Get("abc"); ok {
		fmt.Println(f.(*exampleAbc).Cc)
		f.(*exampleAbc).Cc += 9
	}

	if f, ok := c.Get("abc"); ok {
		fmt.Println(f.(*exampleAbc).Cc)
		//Output:
		//119
	}
}

func ExampleContainer_Factory() {
	c := NewContainer()
	c.Factory("abc", func(cc *Container) interface{} {
		return newExampleAbc()
	})
	if f, ok := c.Get("abc"); ok {
		fmt.Println(f.(*exampleAbc).Cc)
		f.(*exampleAbc).Cc += 9
	}

	if f, ok := c.Get("abc"); ok {
		fmt.Println(f.(*exampleAbc).Cc)
		//Output:
		//110
	}
}