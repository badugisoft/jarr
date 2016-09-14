// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

func EveryString(slice []string, f func(e string) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryStringIndex(slice []string, f func(i int, e string) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryInt(slice []int, f func(e int) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryIntIndex(slice []int, f func(i int, e int) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryUint(slice []uint, f func(e uint) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryUintIndex(slice []uint, f func(i int, e uint) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryInt8(slice []int8, f func(e int8) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryInt8Index(slice []int8, f func(i int, e int8) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryUint8(slice []uint8, f func(e uint8) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryUint8Index(slice []uint8, f func(i int, e uint8) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryInt16(slice []int16, f func(e int16) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryInt16Index(slice []int16, f func(i int, e int16) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryUint16(slice []uint16, f func(e uint16) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryUint16Index(slice []uint16, f func(i int, e uint16) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryInt32(slice []int32, f func(e int32) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryInt32Index(slice []int32, f func(i int, e int32) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryUint32(slice []uint32, f func(e uint32) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryUint32Index(slice []uint32, f func(i int, e uint32) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryInt64(slice []int64, f func(e int64) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryInt64Index(slice []int64, f func(i int, e int64) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryUint64(slice []uint64, f func(e uint64) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryUint64Index(slice []uint64, f func(i int, e uint64) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryBool(slice []bool, f func(e bool) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryBoolIndex(slice []bool, f func(i int, e bool) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}

func EveryRune(slice []rune, f func(e rune) bool) bool {
	for _, e := range slice {
		if !f(e) {
			return false
		}
	}
	return true
}

func EveryRuneIndex(slice []rune, f func(i int, e rune) bool) bool {
	for i, e := range slice {
		if !f(i, e) {
			return false
		}
	}
	return true
}
