package jarr

import "reflect"

func Filter(slice interface{}, f func(e interface{}) bool) interface{} {
	v := reflect.ValueOf(slice)
	ret := reflect.New(reflect.TypeOf(slice)).Elem()
	ret.Set(reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len()))

	for i, j := 0, 0; i < v.Len(); i++ {
		e := v.Index(i)
		if f(e.Interface()) {
			ret.SetLen(j + 1)
			ret.Index(j).Set(e)
			j++
		}
	}
	return ret.Interface()
}

func FilterIndex(slice interface{}, f func(i int, e interface{}) bool) interface{} {
	v := reflect.ValueOf(slice)
	ret := reflect.New(reflect.TypeOf(slice)).Elem()
	ret.Set(reflect.MakeSlice(reflect.TypeOf(slice), 0, v.Len()))

	for i, j := 0, 0; i < v.Len(); i++ {
		e := v.Index(i)
		if f(i, e.Interface()) {
			ret.SetLen(j + 1)
			ret.Index(j).Set(e)
			j++
		}
	}
	return ret.Interface()
}
