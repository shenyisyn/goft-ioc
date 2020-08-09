package Injector

import (
	"reflect"
)

type BeanMapper map[reflect.Type]reflect.Value

func (this BeanMapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)
	if t.Kind() != reflect.Ptr {
		panic("require ptr object")
	}
	this[t] = reflect.ValueOf(bean)
}
func (this BeanMapper) get(bean interface{}) reflect.Value {
	var t reflect.Type
	if bt, ok := bean.(reflect.Type); ok {
		t = bt
	} else {
		t = reflect.TypeOf(bean)
	}
	if v, ok := this[t]; ok {
		return v
	}
	//处理接口 继承
	for k, v := range this {
		if t.Kind() == reflect.Interface && k.Implements(t) {
			return v
		}
	}

	return reflect.Value{}
}
