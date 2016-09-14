// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr

func FilterString(slice []string, f func(e string) bool) []string {
	ret := []string{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterStringIndex(slice []string, f func(i int, e string) bool) []string {
	ret := []string{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt(slice []int, f func(e int) bool) []int {
	ret := []int{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterIntIndex(slice []int, f func(i int, e int) bool) []int {
	ret := []int{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint(slice []uint, f func(e uint) bool) []uint {
	ret := []uint{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUintIndex(slice []uint, f func(i int, e uint) bool) []uint {
	ret := []uint{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt8(slice []int8, f func(e int8) bool) []int8 {
	ret := []int8{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt8Index(slice []int8, f func(i int, e int8) bool) []int8 {
	ret := []int8{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint8(slice []uint8, f func(e uint8) bool) []uint8 {
	ret := []uint8{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint8Index(slice []uint8, f func(i int, e uint8) bool) []uint8 {
	ret := []uint8{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt16(slice []int16, f func(e int16) bool) []int16 {
	ret := []int16{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt16Index(slice []int16, f func(i int, e int16) bool) []int16 {
	ret := []int16{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint16(slice []uint16, f func(e uint16) bool) []uint16 {
	ret := []uint16{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint16Index(slice []uint16, f func(i int, e uint16) bool) []uint16 {
	ret := []uint16{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt32(slice []int32, f func(e int32) bool) []int32 {
	ret := []int32{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt32Index(slice []int32, f func(i int, e int32) bool) []int32 {
	ret := []int32{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint32(slice []uint32, f func(e uint32) bool) []uint32 {
	ret := []uint32{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint32Index(slice []uint32, f func(i int, e uint32) bool) []uint32 {
	ret := []uint32{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt64(slice []int64, f func(e int64) bool) []int64 {
	ret := []int64{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterInt64Index(slice []int64, f func(i int, e int64) bool) []int64 {
	ret := []int64{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint64(slice []uint64, f func(e uint64) bool) []uint64 {
	ret := []uint64{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterUint64Index(slice []uint64, f func(i int, e uint64) bool) []uint64 {
	ret := []uint64{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterBool(slice []bool, f func(e bool) bool) []bool {
	ret := []bool{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterBoolIndex(slice []bool, f func(i int, e bool) bool) []bool {
	ret := []bool{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterRune(slice []rune, f func(e rune) bool) []rune {
	ret := []rune{}
	for _, e := range slice {
		if f(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func FilterRuneIndex(slice []rune, f func(i int, e rune) bool) []rune {
	ret := []rune{}
	for i, e := range slice {
		if f(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}
