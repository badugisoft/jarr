package jarr

import "reflect"

func Each(slice interface{}, f func(e interface{})) interface{} {
	v := reflect.ValueOf(slice)
	for i := 0; i < v.Len(); i++ {
		f(v.Index(i).Addr().Interface())
	}
	return slice
}

func EachIndex(slice interface{}, f func(i int, e interface{})) interface{} {
	v := reflect.ValueOf(slice)
	for i := 0; i < v.Len(); i++ {
		f(i, v.Index(i).Addr().Interface())
	}
	return slice
}
