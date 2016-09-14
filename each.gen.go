// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

func EachString(slice []string, f func(e *string)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachStringIndex(slice []string, f func(i int, e *string)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachInt(slice []int, f func(e *int)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachIntIndex(slice []int, f func(i int, e *int)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachUint(slice []uint, f func(e *uint)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachUintIndex(slice []uint, f func(i int, e *uint)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachInt8(slice []int8, f func(e *int8)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachInt8Index(slice []int8, f func(i int, e *int8)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachUint8(slice []uint8, f func(e *uint8)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachUint8Index(slice []uint8, f func(i int, e *uint8)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachInt16(slice []int16, f func(e *int16)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachInt16Index(slice []int16, f func(i int, e *int16)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachUint16(slice []uint16, f func(e *uint16)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachUint16Index(slice []uint16, f func(i int, e *uint16)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachInt32(slice []int32, f func(e *int32)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachInt32Index(slice []int32, f func(i int, e *int32)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachUint32(slice []uint32, f func(e *uint32)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachUint32Index(slice []uint32, f func(i int, e *uint32)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachInt64(slice []int64, f func(e *int64)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachInt64Index(slice []int64, f func(i int, e *int64)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachUint64(slice []uint64, f func(e *uint64)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachUint64Index(slice []uint64, f func(i int, e *uint64)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachBool(slice []bool, f func(e *bool)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachBoolIndex(slice []bool, f func(i int, e *bool)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}

func EachRune(slice []rune, f func(e *rune)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(&slice[i])
	}
}

func EachRuneIndex(slice []rune, f func(i int, e *rune)) {
	for i, end := 0, len(slice); i < end; i++ {
		f(i, &slice[i])
	}
}
