// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

func CondString(c bool, t, f string) string {
	if c {
		return t
	} else {
		return f
	}
}

func CondInt(c bool, t, f int) int {
	if c {
		return t
	} else {
		return f
	}
}

func CondUint(c bool, t, f uint) uint {
	if c {
		return t
	} else {
		return f
	}
}

func CondInt8(c bool, t, f int8) int8 {
	if c {
		return t
	} else {
		return f
	}
}

func CondUint8(c bool, t, f uint8) uint8 {
	if c {
		return t
	} else {
		return f
	}
}

func CondInt16(c bool, t, f int16) int16 {
	if c {
		return t
	} else {
		return f
	}
}

func CondUint16(c bool, t, f uint16) uint16 {
	if c {
		return t
	} else {
		return f
	}
}

func CondInt32(c bool, t, f int32) int32 {
	if c {
		return t
	} else {
		return f
	}
}

func CondUint32(c bool, t, f uint32) uint32 {
	if c {
		return t
	} else {
		return f
	}
}

func CondInt64(c bool, t, f int64) int64 {
	if c {
		return t
	} else {
		return f
	}
}

func CondUint64(c bool, t, f uint64) uint64 {
	if c {
		return t
	} else {
		return f
	}
}

func CondBool(c bool, t, f bool) bool {
	if c {
		return t
	} else {
		return f
	}
}

func CondRune(c bool, t, f rune) rune {
	if c {
		return t
	} else {
		return f
	}
}

func ReverseString(slice []string) []string {
	l := len(slice)
	r := make([]string, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseInt(slice []int) []int {
	l := len(slice)
	r := make([]int, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseUint(slice []uint) []uint {
	l := len(slice)
	r := make([]uint, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseInt8(slice []int8) []int8 {
	l := len(slice)
	r := make([]int8, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseUint8(slice []uint8) []uint8 {
	l := len(slice)
	r := make([]uint8, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseInt16(slice []int16) []int16 {
	l := len(slice)
	r := make([]int16, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseUint16(slice []uint16) []uint16 {
	l := len(slice)
	r := make([]uint16, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseInt32(slice []int32) []int32 {
	l := len(slice)
	r := make([]int32, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseUint32(slice []uint32) []uint32 {
	l := len(slice)
	r := make([]uint32, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseInt64(slice []int64) []int64 {
	l := len(slice)
	r := make([]int64, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseUint64(slice []uint64) []uint64 {
	l := len(slice)
	r := make([]uint64, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseBool(slice []bool) []bool {
	l := len(slice)
	r := make([]bool, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}

func ReverseRune(slice []rune) []rune {
	l := len(slice)
	r := make([]rune, l)
	for i, e := range slice {
		r[l-i-1] = e
	}
	return r
}
