// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

import "reflect"

func ReduceToString(slice interface{}, f func(prev string, e interface{}) string, init string) string {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToStringIndex(slice interface{}, f func(prev string, i int, e interface{}) string, init string) string {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceString(slice []string, f func(prev string, e string) string, init string) string {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceStringIndex(slice []string, f func(prev string, i int, e string) string, init string) string {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToInt(slice interface{}, f func(prev int, e interface{}) int, init int) int {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToIntIndex(slice interface{}, f func(prev int, i int, e interface{}) int, init int) int {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceInt(slice []int, f func(prev int, e int) int, init int) int {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceIntIndex(slice []int, f func(prev int, i int, e int) int, init int) int {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToUint(slice interface{}, f func(prev uint, e interface{}) uint, init uint) uint {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToUintIndex(slice interface{}, f func(prev uint, i int, e interface{}) uint, init uint) uint {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceUint(slice []uint, f func(prev uint, e uint) uint, init uint) uint {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceUintIndex(slice []uint, f func(prev uint, i int, e uint) uint, init uint) uint {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToInt8(slice interface{}, f func(prev int8, e interface{}) int8, init int8) int8 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToInt8Index(slice interface{}, f func(prev int8, i int, e interface{}) int8, init int8) int8 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceInt8(slice []int8, f func(prev int8, e int8) int8, init int8) int8 {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceInt8Index(slice []int8, f func(prev int8, i int, e int8) int8, init int8) int8 {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToUint8(slice interface{}, f func(prev uint8, e interface{}) uint8, init uint8) uint8 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToUint8Index(slice interface{}, f func(prev uint8, i int, e interface{}) uint8, init uint8) uint8 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceUint8(slice []uint8, f func(prev uint8, e uint8) uint8, init uint8) uint8 {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceUint8Index(slice []uint8, f func(prev uint8, i int, e uint8) uint8, init uint8) uint8 {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToInt16(slice interface{}, f func(prev int16, e interface{}) int16, init int16) int16 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToInt16Index(slice interface{}, f func(prev int16, i int, e interface{}) int16, init int16) int16 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceInt16(slice []int16, f func(prev int16, e int16) int16, init int16) int16 {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceInt16Index(slice []int16, f func(prev int16, i int, e int16) int16, init int16) int16 {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToUint16(slice interface{}, f func(prev uint16, e interface{}) uint16, init uint16) uint16 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToUint16Index(slice interface{}, f func(prev uint16, i int, e interface{}) uint16, init uint16) uint16 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceUint16(slice []uint16, f func(prev uint16, e uint16) uint16, init uint16) uint16 {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceUint16Index(slice []uint16, f func(prev uint16, i int, e uint16) uint16, init uint16) uint16 {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToInt32(slice interface{}, f func(prev int32, e interface{}) int32, init int32) int32 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToInt32Index(slice interface{}, f func(prev int32, i int, e interface{}) int32, init int32) int32 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceInt32(slice []int32, f func(prev int32, e int32) int32, init int32) int32 {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceInt32Index(slice []int32, f func(prev int32, i int, e int32) int32, init int32) int32 {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToUint32(slice interface{}, f func(prev uint32, e interface{}) uint32, init uint32) uint32 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToUint32Index(slice interface{}, f func(prev uint32, i int, e interface{}) uint32, init uint32) uint32 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceUint32(slice []uint32, f func(prev uint32, e uint32) uint32, init uint32) uint32 {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceUint32Index(slice []uint32, f func(prev uint32, i int, e uint32) uint32, init uint32) uint32 {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToInt64(slice interface{}, f func(prev int64, e interface{}) int64, init int64) int64 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToInt64Index(slice interface{}, f func(prev int64, i int, e interface{}) int64, init int64) int64 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceInt64(slice []int64, f func(prev int64, e int64) int64, init int64) int64 {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceInt64Index(slice []int64, f func(prev int64, i int, e int64) int64, init int64) int64 {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToUint64(slice interface{}, f func(prev uint64, e interface{}) uint64, init uint64) uint64 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToUint64Index(slice interface{}, f func(prev uint64, i int, e interface{}) uint64, init uint64) uint64 {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceUint64(slice []uint64, f func(prev uint64, e uint64) uint64, init uint64) uint64 {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceUint64Index(slice []uint64, f func(prev uint64, i int, e uint64) uint64, init uint64) uint64 {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToBool(slice interface{}, f func(prev bool, e interface{}) bool, init bool) bool {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToBoolIndex(slice interface{}, f func(prev bool, i int, e interface{}) bool, init bool) bool {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceBool(slice []bool, f func(prev bool, e bool) bool, init bool) bool {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceBoolIndex(slice []bool, f func(prev bool, i int, e bool) bool, init bool) bool {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}

func ReduceToRune(slice interface{}, f func(prev rune, e interface{}) rune, init rune) rune {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, v.Index(i).Interface())
	}
	return init
}

func ReduceToRuneIndex(slice interface{}, f func(prev rune, i int, e interface{}) rune, init rune) rune {
	v := reflect.ValueOf(slice)
	for i, end := 0, v.Len(); i < end; i++ {
		init = f(init, i, v.Index(i).Interface())
	}
	return init
}

func ReduceRune(slice []rune, f func(prev rune, e rune) rune, init rune) rune {
	for _, e := range slice {
		init = f(init, e)
	}
	return init
}

func ReduceRuneIndex(slice []rune, f func(prev rune, i int, e rune) rune, init rune) rune {
	for i, e := range slice {
		init = f(init, i, e)
	}
	return init
}
