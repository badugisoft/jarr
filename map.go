// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

import "reflect"

func MapString(slice interface{}, f func(e interface{}) string) []string {
	v := reflect.ValueOf(slice)
	ret := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapStringIndex(slice interface{}, f func(i int, e interface{}) string) []string {
	v := reflect.ValueOf(slice)
	ret := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt(slice interface{}, f func(e interface{}) int) []int {
	v := reflect.ValueOf(slice)
	ret := make([]int, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapIntIndex(slice interface{}, f func(i int, e interface{}) int) []int {
	v := reflect.ValueOf(slice)
	ret := make([]int, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint(slice interface{}, f func(e interface{}) uint) []uint {
	v := reflect.ValueOf(slice)
	ret := make([]uint, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUintIndex(slice interface{}, f func(i int, e interface{}) uint) []uint {
	v := reflect.ValueOf(slice)
	ret := make([]uint, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt8(slice interface{}, f func(e interface{}) int8) []int8 {
	v := reflect.ValueOf(slice)
	ret := make([]int8, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt8Index(slice interface{}, f func(i int, e interface{}) int8) []int8 {
	v := reflect.ValueOf(slice)
	ret := make([]int8, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint8(slice interface{}, f func(e interface{}) uint8) []uint8 {
	v := reflect.ValueOf(slice)
	ret := make([]uint8, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint8Index(slice interface{}, f func(i int, e interface{}) uint8) []uint8 {
	v := reflect.ValueOf(slice)
	ret := make([]uint8, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt16(slice interface{}, f func(e interface{}) int16) []int16 {
	v := reflect.ValueOf(slice)
	ret := make([]int16, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt16Index(slice interface{}, f func(i int, e interface{}) int16) []int16 {
	v := reflect.ValueOf(slice)
	ret := make([]int16, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint16(slice interface{}, f func(e interface{}) uint16) []uint16 {
	v := reflect.ValueOf(slice)
	ret := make([]uint16, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint16Index(slice interface{}, f func(i int, e interface{}) uint16) []uint16 {
	v := reflect.ValueOf(slice)
	ret := make([]uint16, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt32(slice interface{}, f func(e interface{}) int32) []int32 {
	v := reflect.ValueOf(slice)
	ret := make([]int32, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt32Index(slice interface{}, f func(i int, e interface{}) int32) []int32 {
	v := reflect.ValueOf(slice)
	ret := make([]int32, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint32(slice interface{}, f func(e interface{}) uint32) []uint32 {
	v := reflect.ValueOf(slice)
	ret := make([]uint32, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint32Index(slice interface{}, f func(i int, e interface{}) uint32) []uint32 {
	v := reflect.ValueOf(slice)
	ret := make([]uint32, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt64(slice interface{}, f func(e interface{}) int64) []int64 {
	v := reflect.ValueOf(slice)
	ret := make([]int64, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapInt64Index(slice interface{}, f func(i int, e interface{}) int64) []int64 {
	v := reflect.ValueOf(slice)
	ret := make([]int64, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint64(slice interface{}, f func(e interface{}) uint64) []uint64 {
	v := reflect.ValueOf(slice)
	ret := make([]uint64, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapUint64Index(slice interface{}, f func(i int, e interface{}) uint64) []uint64 {
	v := reflect.ValueOf(slice)
	ret := make([]uint64, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapBool(slice interface{}, f func(e interface{}) bool) []bool {
	v := reflect.ValueOf(slice)
	ret := make([]bool, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapBoolIndex(slice interface{}, f func(i int, e interface{}) bool) []bool {
	v := reflect.ValueOf(slice)
	ret := make([]bool, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

func MapRune(slice interface{}, f func(e interface{}) rune) []rune {
	v := reflect.ValueOf(slice)
	ret := make([]rune, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

func MapRuneIndex(slice interface{}, f func(i int, e interface{}) rune) []rune {
	v := reflect.ValueOf(slice)
	ret := make([]rune, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}
