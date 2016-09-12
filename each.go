// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

func EachString(s []string, f func(e *string)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachStringIndex(s []string, f func(i int, e *string)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachInt(s []int, f func(e *int)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachIntIndex(s []int, f func(i int, e *int)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachUint(s []uint, f func(e *uint)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachUintIndex(s []uint, f func(i int, e *uint)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachInt8(s []int8, f func(e *int8)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachInt8Index(s []int8, f func(i int, e *int8)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachUint8(s []uint8, f func(e *uint8)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachUint8Index(s []uint8, f func(i int, e *uint8)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachInt16(s []int16, f func(e *int16)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachInt16Index(s []int16, f func(i int, e *int16)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachUint16(s []uint16, f func(e *uint16)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachUint16Index(s []uint16, f func(i int, e *uint16)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachInt32(s []int32, f func(e *int32)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachInt32Index(s []int32, f func(i int, e *int32)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachUint32(s []uint32, f func(e *uint32)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachUint32Index(s []uint32, f func(i int, e *uint32)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachInt64(s []int64, f func(e *int64)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachInt64Index(s []int64, f func(i int, e *int64)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachUint64(s []uint64, f func(e *uint64)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachUint64Index(s []uint64, f func(i int, e *uint64)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func EachBool(s []bool, f func(e *bool)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachBoolIndex(s []bool, f func(i int, e *bool)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}

func Eachrune(s []rune, f func(e *rune)) {
	for i, l := 0, len(s); i < l; i++ {
		f(&s[i])
	}
}

func EachruneIndex(s []rune, f func(i int, e *rune)) {
	for i, l := 0, len(s); i < l; i++ {
		f(i, &s[i])
	}
}
