// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

import "reflect"

type MapStringFunc func(e interface{}) string

func MapString(slice interface{}, f MapStringFunc) []string {
	v := reflect.ValueOf(slice)
	ret := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapStringIndexFunc func(i int, e interface{}) string

func MapStringIndex(slice interface{}, f MapStringIndexFunc) []string {
	v := reflect.ValueOf(slice)
	ret := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapIntFunc func(e interface{}) int

func MapInt(slice interface{}, f MapIntFunc) []int {
	v := reflect.ValueOf(slice)
	ret := make([]int, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapIntIndexFunc func(i int, e interface{}) int

func MapIntIndex(slice interface{}, f MapIntIndexFunc) []int {
	v := reflect.ValueOf(slice)
	ret := make([]int, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUintFunc func(e interface{}) uint

func MapUint(slice interface{}, f MapUintFunc) []uint {
	v := reflect.ValueOf(slice)
	ret := make([]uint, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUintIndexFunc func(i int, e interface{}) uint

func MapUintIndex(slice interface{}, f MapUintIndexFunc) []uint {
	v := reflect.ValueOf(slice)
	ret := make([]uint, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapInt8Func func(e interface{}) int8

func MapInt8(slice interface{}, f MapInt8Func) []int8 {
	v := reflect.ValueOf(slice)
	ret := make([]int8, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapInt8IndexFunc func(i int, e interface{}) int8

func MapInt8Index(slice interface{}, f MapInt8IndexFunc) []int8 {
	v := reflect.ValueOf(slice)
	ret := make([]int8, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUint8Func func(e interface{}) uint8

func MapUint8(slice interface{}, f MapUint8Func) []uint8 {
	v := reflect.ValueOf(slice)
	ret := make([]uint8, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUint8IndexFunc func(i int, e interface{}) uint8

func MapUint8Index(slice interface{}, f MapUint8IndexFunc) []uint8 {
	v := reflect.ValueOf(slice)
	ret := make([]uint8, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapInt16Func func(e interface{}) int16

func MapInt16(slice interface{}, f MapInt16Func) []int16 {
	v := reflect.ValueOf(slice)
	ret := make([]int16, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapInt16IndexFunc func(i int, e interface{}) int16

func MapInt16Index(slice interface{}, f MapInt16IndexFunc) []int16 {
	v := reflect.ValueOf(slice)
	ret := make([]int16, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUint16Func func(e interface{}) uint16

func MapUint16(slice interface{}, f MapUint16Func) []uint16 {
	v := reflect.ValueOf(slice)
	ret := make([]uint16, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUint16IndexFunc func(i int, e interface{}) uint16

func MapUint16Index(slice interface{}, f MapUint16IndexFunc) []uint16 {
	v := reflect.ValueOf(slice)
	ret := make([]uint16, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapInt32Func func(e interface{}) int32

func MapInt32(slice interface{}, f MapInt32Func) []int32 {
	v := reflect.ValueOf(slice)
	ret := make([]int32, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapInt32IndexFunc func(i int, e interface{}) int32

func MapInt32Index(slice interface{}, f MapInt32IndexFunc) []int32 {
	v := reflect.ValueOf(slice)
	ret := make([]int32, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUint32Func func(e interface{}) uint32

func MapUint32(slice interface{}, f MapUint32Func) []uint32 {
	v := reflect.ValueOf(slice)
	ret := make([]uint32, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUint32IndexFunc func(i int, e interface{}) uint32

func MapUint32Index(slice interface{}, f MapUint32IndexFunc) []uint32 {
	v := reflect.ValueOf(slice)
	ret := make([]uint32, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapInt64Func func(e interface{}) int64

func MapInt64(slice interface{}, f MapInt64Func) []int64 {
	v := reflect.ValueOf(slice)
	ret := make([]int64, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapInt64IndexFunc func(i int, e interface{}) int64

func MapInt64Index(slice interface{}, f MapInt64IndexFunc) []int64 {
	v := reflect.ValueOf(slice)
	ret := make([]int64, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUint64Func func(e interface{}) uint64

func MapUint64(slice interface{}, f MapUint64Func) []uint64 {
	v := reflect.ValueOf(slice)
	ret := make([]uint64, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapUint64IndexFunc func(i int, e interface{}) uint64

func MapUint64Index(slice interface{}, f MapUint64IndexFunc) []uint64 {
	v := reflect.ValueOf(slice)
	ret := make([]uint64, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapBoolFunc func(e interface{}) bool

func MapBool(slice interface{}, f MapBoolFunc) []bool {
	v := reflect.ValueOf(slice)
	ret := make([]bool, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapBoolIndexFunc func(i int, e interface{}) bool

func MapBoolIndex(slice interface{}, f MapBoolIndexFunc) []bool {
	v := reflect.ValueOf(slice)
	ret := make([]bool, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}

type MapruneFunc func(e interface{}) rune

func Maprune(slice interface{}, f MapruneFunc) []rune {
	v := reflect.ValueOf(slice)
	ret := make([]rune, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(v.Index(i).Addr().Interface())
	}
	return ret
}

type MapruneIndexFunc func(i int, e interface{}) rune

func MapruneIndex(slice interface{}, f MapruneIndexFunc) []rune {
	v := reflect.ValueOf(slice)
	ret := make([]rune, v.Len())
	for i := 0; i < v.Len(); i++ {
		ret[i] = f(i, v.Index(i).Addr().Interface())
	}
	return ret
}
