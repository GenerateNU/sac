package cache

import (
	"reflect"
)

func SetPointedValue(dest interface{}, src interface{}) {
	reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(src).Elem())
}
