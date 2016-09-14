package jarr

import "reflect"

func Map(slice interface{}, f func(e interface{}) interface{}) interface{} {
	v := reflect.ValueOf(slice)
	ret := reflect.MakeSlice(reflect.TypeOf(slice), v.Len(), v.Len())
	for i := 0; i < v.Len(); i++ {
		ret.Index(i).Set(reflect.ValueOf(f(v.Index(i).Interface())))
	}
	return ret.Interface()
}

func MapIndex(slice interface{}, f func(i int, e interface{}) interface{}) interface{} {
	v := reflect.ValueOf(slice)
	ret := reflect.MakeSlice(reflect.TypeOf(slice), v.Len(), v.Len())
	for i := 0; i < v.Len(); i++ {
		ret.Index(i).Set(reflect.ValueOf(f(i, v.Index(i).Interface())))
	}
	return ret.Interface()
}

func MapType(slice interface{}, t interface{}, f func(e interface{}) interface{}) interface{} {
	v := reflect.ValueOf(slice)
	ret := reflect.MakeSlice(reflect.TypeOf(t), v.Len(), v.Len())
	for i := 0; i < v.Len(); i++ {
		ret.Index(i).Set(reflect.ValueOf(f(v.Index(i).Interface())))
	}
	return ret.Interface()
}

func MapTypeIndex(slice interface{}, t interface{}, f func(i int, e interface{}) interface{}) interface{} {
	v := reflect.ValueOf(slice)
	ret := reflect.MakeSlice(reflect.TypeOf(t), v.Len(), v.Len())
	for i := 0; i < v.Len(); i++ {
		ret.Index(i).Set(reflect.ValueOf(f(i, v.Index(i).Interface())))
	}
	return ret.Interface()
}
