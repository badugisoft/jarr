// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

import "reflect"

func MapString(slice []string, f func(e string) string) []string {
	ret := make([]string, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapStringIndex(slice []string, f func(i int, e string) string) []string {
	ret := make([]string, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToString(slice interface{}, f func(e interface{}) string) []string {
	v := reflect.ValueOf(slice)
	ret := make([]string, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToStringIndex(slice interface{}, f func(i int, e interface{}) string) []string {
	v := reflect.ValueOf(slice)
	ret := make([]string, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapInt(slice []int, f func(e int) int) []int {
	ret := make([]int, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapIntIndex(slice []int, f func(i int, e int) int) []int {
	ret := make([]int, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToInt(slice interface{}, f func(e interface{}) int) []int {
	v := reflect.ValueOf(slice)
	ret := make([]int, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToIntIndex(slice interface{}, f func(i int, e interface{}) int) []int {
	v := reflect.ValueOf(slice)
	ret := make([]int, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapUint(slice []uint, f func(e uint) uint) []uint {
	ret := make([]uint, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapUintIndex(slice []uint, f func(i int, e uint) uint) []uint {
	ret := make([]uint, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToUint(slice interface{}, f func(e interface{}) uint) []uint {
	v := reflect.ValueOf(slice)
	ret := make([]uint, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToUintIndex(slice interface{}, f func(i int, e interface{}) uint) []uint {
	v := reflect.ValueOf(slice)
	ret := make([]uint, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapInt8(slice []int8, f func(e int8) int8) []int8 {
	ret := make([]int8, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapInt8Index(slice []int8, f func(i int, e int8) int8) []int8 {
	ret := make([]int8, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToInt8(slice interface{}, f func(e interface{}) int8) []int8 {
	v := reflect.ValueOf(slice)
	ret := make([]int8, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToInt8Index(slice interface{}, f func(i int, e interface{}) int8) []int8 {
	v := reflect.ValueOf(slice)
	ret := make([]int8, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapUint8(slice []uint8, f func(e uint8) uint8) []uint8 {
	ret := make([]uint8, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapUint8Index(slice []uint8, f func(i int, e uint8) uint8) []uint8 {
	ret := make([]uint8, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToUint8(slice interface{}, f func(e interface{}) uint8) []uint8 {
	v := reflect.ValueOf(slice)
	ret := make([]uint8, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToUint8Index(slice interface{}, f func(i int, e interface{}) uint8) []uint8 {
	v := reflect.ValueOf(slice)
	ret := make([]uint8, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapInt16(slice []int16, f func(e int16) int16) []int16 {
	ret := make([]int16, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapInt16Index(slice []int16, f func(i int, e int16) int16) []int16 {
	ret := make([]int16, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToInt16(slice interface{}, f func(e interface{}) int16) []int16 {
	v := reflect.ValueOf(slice)
	ret := make([]int16, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToInt16Index(slice interface{}, f func(i int, e interface{}) int16) []int16 {
	v := reflect.ValueOf(slice)
	ret := make([]int16, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapUint16(slice []uint16, f func(e uint16) uint16) []uint16 {
	ret := make([]uint16, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapUint16Index(slice []uint16, f func(i int, e uint16) uint16) []uint16 {
	ret := make([]uint16, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToUint16(slice interface{}, f func(e interface{}) uint16) []uint16 {
	v := reflect.ValueOf(slice)
	ret := make([]uint16, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToUint16Index(slice interface{}, f func(i int, e interface{}) uint16) []uint16 {
	v := reflect.ValueOf(slice)
	ret := make([]uint16, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapInt32(slice []int32, f func(e int32) int32) []int32 {
	ret := make([]int32, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapInt32Index(slice []int32, f func(i int, e int32) int32) []int32 {
	ret := make([]int32, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToInt32(slice interface{}, f func(e interface{}) int32) []int32 {
	v := reflect.ValueOf(slice)
	ret := make([]int32, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToInt32Index(slice interface{}, f func(i int, e interface{}) int32) []int32 {
	v := reflect.ValueOf(slice)
	ret := make([]int32, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapUint32(slice []uint32, f func(e uint32) uint32) []uint32 {
	ret := make([]uint32, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapUint32Index(slice []uint32, f func(i int, e uint32) uint32) []uint32 {
	ret := make([]uint32, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToUint32(slice interface{}, f func(e interface{}) uint32) []uint32 {
	v := reflect.ValueOf(slice)
	ret := make([]uint32, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToUint32Index(slice interface{}, f func(i int, e interface{}) uint32) []uint32 {
	v := reflect.ValueOf(slice)
	ret := make([]uint32, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapInt64(slice []int64, f func(e int64) int64) []int64 {
	ret := make([]int64, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapInt64Index(slice []int64, f func(i int, e int64) int64) []int64 {
	ret := make([]int64, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToInt64(slice interface{}, f func(e interface{}) int64) []int64 {
	v := reflect.ValueOf(slice)
	ret := make([]int64, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToInt64Index(slice interface{}, f func(i int, e interface{}) int64) []int64 {
	v := reflect.ValueOf(slice)
	ret := make([]int64, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapUint64(slice []uint64, f func(e uint64) uint64) []uint64 {
	ret := make([]uint64, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapUint64Index(slice []uint64, f func(i int, e uint64) uint64) []uint64 {
	ret := make([]uint64, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToUint64(slice interface{}, f func(e interface{}) uint64) []uint64 {
	v := reflect.ValueOf(slice)
	ret := make([]uint64, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToUint64Index(slice interface{}, f func(i int, e interface{}) uint64) []uint64 {
	v := reflect.ValueOf(slice)
	ret := make([]uint64, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapBool(slice []bool, f func(e bool) bool) []bool {
	ret := make([]bool, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapBoolIndex(slice []bool, f func(i int, e bool) bool) []bool {
	ret := make([]bool, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToBool(slice interface{}, f func(e interface{}) bool) []bool {
	v := reflect.ValueOf(slice)
	ret := make([]bool, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToBoolIndex(slice interface{}, f func(i int, e interface{}) bool) []bool {
	v := reflect.ValueOf(slice)
	ret := make([]bool, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}

func MapRune(slice []rune, f func(e rune) rune) []rune {
	ret := make([]rune, len(slice))
	for i, e := range slice {
		ret[i] = f(e)
	}
	return ret
}

func MapRuneIndex(slice []rune, f func(i int, e rune) rune) []rune {
	ret := make([]rune, len(slice))
	for i, e := range slice {
		ret[i] = f(i, e)
	}
	return ret
}

func MapToRune(slice interface{}, f func(e interface{}) rune) []rune {
	v := reflect.ValueOf(slice)
	ret := make([]rune, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(v.Index(i).Interface())
	}
	return ret
}

func MapToRuneIndex(slice interface{}, f func(i int, e interface{}) rune) []rune {
	v := reflect.ValueOf(slice)
	ret := make([]rune, v.Len())
	for i, end := 0, v.Len(); i < end; i++ {
		ret[i] = f(i, v.Index(i).Interface())
	}
	return ret
}
