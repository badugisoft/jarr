package jarr

import "reflect"

func Every(slice interface{}, f func(e interface{}) bool) bool {
	v := reflect.ValueOf(slice)
	for i := 0; i < v.Len(); i++ {
		if !f(v.Index(i).Interface()) {
			return false
		}
	}
	return true
}

func EveryIndex(slice interface{}, f func(i int, e interface{}) bool) bool {
	v := reflect.ValueOf(slice)
	for i := 0; i < v.Len(); i++ {
		if !f(i, v.Index(i).Interface()) {
			return false
		}
	}
	return true
}
