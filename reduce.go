package jarr

import "reflect"

func Reduce(slice interface{}, f func(prev interface{}, e interface{}) interface{}, init interface{}) interface{} {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceIndex(slice interface{}, f func(prev interface{}, i int, e interface{}) interface{}, init interface{}) interface{} {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}
