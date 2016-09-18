// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

func EachString(slice []string, f func(e *string)) []string {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachStringIndex(slice []string, f func(i int, e *string)) []string {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachInt(slice []int, f func(e *int)) []int {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachIntIndex(slice []int, f func(i int, e *int)) []int {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachUint(slice []uint, f func(e *uint)) []uint {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachUintIndex(slice []uint, f func(i int, e *uint)) []uint {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachInt8(slice []int8, f func(e *int8)) []int8 {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachInt8Index(slice []int8, f func(i int, e *int8)) []int8 {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachUint8(slice []uint8, f func(e *uint8)) []uint8 {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachUint8Index(slice []uint8, f func(i int, e *uint8)) []uint8 {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachInt16(slice []int16, f func(e *int16)) []int16 {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachInt16Index(slice []int16, f func(i int, e *int16)) []int16 {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachUint16(slice []uint16, f func(e *uint16)) []uint16 {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachUint16Index(slice []uint16, f func(i int, e *uint16)) []uint16 {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachInt32(slice []int32, f func(e *int32)) []int32 {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachInt32Index(slice []int32, f func(i int, e *int32)) []int32 {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachUint32(slice []uint32, f func(e *uint32)) []uint32 {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachUint32Index(slice []uint32, f func(i int, e *uint32)) []uint32 {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachInt64(slice []int64, f func(e *int64)) []int64 {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachInt64Index(slice []int64, f func(i int, e *int64)) []int64 {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachUint64(slice []uint64, f func(e *uint64)) []uint64 {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachUint64Index(slice []uint64, f func(i int, e *uint64)) []uint64 {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachBool(slice []bool, f func(e *bool)) []bool {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachBoolIndex(slice []bool, f func(i int, e *bool)) []bool {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}

func EachRune(slice []rune, f func(e *rune)) []rune {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
	return slice
}

func EachRuneIndex(slice []rune, f func(i int, e *rune)) []rune {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
	return slice
}
