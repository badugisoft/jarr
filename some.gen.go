// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

func SomeString(slice []string, f func(e string) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeStringIndex(slice []string, f func(i int, e string) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeInt(slice []int, f func(e int) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeIntIndex(slice []int, f func(i int, e int) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeUint(slice []uint, f func(e uint) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeUintIndex(slice []uint, f func(i int, e uint) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeInt8(slice []int8, f func(e int8) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeInt8Index(slice []int8, f func(i int, e int8) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeUint8(slice []uint8, f func(e uint8) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeUint8Index(slice []uint8, f func(i int, e uint8) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeInt16(slice []int16, f func(e int16) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeInt16Index(slice []int16, f func(i int, e int16) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeUint16(slice []uint16, f func(e uint16) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeUint16Index(slice []uint16, f func(i int, e uint16) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeInt32(slice []int32, f func(e int32) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeInt32Index(slice []int32, f func(i int, e int32) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeUint32(slice []uint32, f func(e uint32) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeUint32Index(slice []uint32, f func(i int, e uint32) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeInt64(slice []int64, f func(e int64) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeInt64Index(slice []int64, f func(i int, e int64) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeUint64(slice []uint64, f func(e uint64) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeUint64Index(slice []uint64, f func(i int, e uint64) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeBool(slice []bool, f func(e bool) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeBoolIndex(slice []bool, f func(i int, e bool) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}

func SomeRune(slice []rune, f func(e rune) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func SomeRuneIndex(slice []rune, f func(i int, e rune) bool) bool {
	for i, e := range slice {
		if f(i, e) {
			return true
		}
	}
	return false
}
