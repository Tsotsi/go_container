package go_container

import (
	"sync"
	"reflect"
)

//
//
//

type ItemType int

const (
	ITEM_TYPE_SHARED ItemType = 1 << iota
	ITEM_TYPE_FACTORY
	ITEM_TYPE_VALUE
)

type Container struct {
	container *sync.Map
}

type Item struct {
	Type    ItemType
	one     *sync.Once
	Raw     interface{}
	Value   interface{}
	rawInfo [2]int
}

func NewContainer() *Container {
	c := new(Container)
	c.container = new(sync.Map)
	return c
}

// Shared Value
func (c *Container) Set(key, value interface{}) {
	t := reflect.TypeOf(value)
	v := new(Item)
	v.Type = ITEM_TYPE_VALUE
	v.Raw = value
	//func
	if t.Kind() == reflect.Func {
		if t.NumOut() != 1 {
			panic("func return too much")
		}
		if t.NumIn() > 1 {
			panic("func argc need under 1")
		}
		v.Type = ITEM_TYPE_SHARED
		v.rawInfo[0] = t.NumIn()
		v.rawInfo[1] = t.NumOut()
		v.one = new(sync.Once)
	}

	c.container.Store(key, v)

}

func (c Container) Get(key interface{}) (value interface{}, ok bool) {
	//first factory
	if vc, exists := c.container.Load(key); exists {
		ok = exists
		v := vc.(*Item)
		switch v.Type {
		case ITEM_TYPE_VALUE:
			value = v.Raw
		case ITEM_TYPE_SHARED:
			v.one.Do(func() {
				vv := reflect.ValueOf(v.Raw)
				var args []reflect.Value
				if v.rawInfo[0] > 0 {
					args = append(args, reflect.ValueOf(&c))
				}
				resv := vv.Call(args)
				v.Value = resv[0].Interface()
			})
			value = v.Value

		case ITEM_TYPE_FACTORY:
			vv := reflect.ValueOf(v.Raw)
			args := []reflect.Value{reflect.ValueOf(&c)}
			resv := vv.Call(args)
			v.Value = resv[0].Interface()
			value = v.Value
		}
	}
	return
}

func (c *Container) Factory(key interface{}, value func(container *Container) interface{}) {
	v := new(Item)
	v.Type = ITEM_TYPE_FACTORY
	v.Raw = value
	c.container.Store(key, v)
}

func (c *Container) Raw(key interface{}) (value interface{}, ok bool) {
	var v interface{}
	if v, ok = c.container.Load(key); ok {
		value = v.(*Item).Raw
	}
	return
}
