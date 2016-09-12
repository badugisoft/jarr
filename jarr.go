package jarr

import "reflect"

func Equal(s1, s2 interface{}) bool {
	return reflect.DeepEqual(s1, s2)
}

func Each(slice interface{}, f func(e interface{})) {
	v := reflect.ValueOf(slice)
	for i := 0; i < v.Len(); i++ {
		f(v.Index(i).Addr().Interface())
	}
}

func EachIndex(slice interface{}, f func(i int, e interface{})) {
	v := reflect.ValueOf(slice)
	for i := 0; i < v.Len(); i++ {
		f(i, v.Index(i).Addr().Interface())
	}
}
