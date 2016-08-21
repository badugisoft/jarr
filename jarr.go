// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package jarr


type String []string

func (s String) Equal(v String) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type StringMapFunc func(current string, index int, array []string) string

func (s String) Map(f StringMapFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type StringReduceFunc func(previousValue string, currentValue string, currentIndex int, array []string) string

func (s String) Reduce(f StringReduceFunc, initialValue string) string {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type StringTestFunc func(current string, index int, array []string) bool

func (s String) Some(f StringTestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s String) Every(f StringTestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s String) Filter(f StringTestFunc) String {
	r := []string{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Bool []bool

func (s Bool) Equal(v Bool) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type BoolMapFunc func(current bool, index int, array []bool) bool

func (s Bool) Map(f BoolMapFunc) Bool {
	r := make([]bool, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type BoolReduceFunc func(previousValue bool, currentValue bool, currentIndex int, array []bool) bool

func (s Bool) Reduce(f BoolReduceFunc, initialValue bool) bool {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type BoolTestFunc func(current bool, index int, array []bool) bool

func (s Bool) Some(f BoolTestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Bool) Every(f BoolTestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Bool) Filter(f BoolTestFunc) Bool {
	r := []bool{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Int []int

func (s Int) Equal(v Int) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type IntMapFunc func(current int, index int, array []int) int

func (s Int) Map(f IntMapFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type IntReduceFunc func(previousValue int, currentValue int, currentIndex int, array []int) int

func (s Int) Reduce(f IntReduceFunc, initialValue int) int {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type IntTestFunc func(current int, index int, array []int) bool

func (s Int) Some(f IntTestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Int) Every(f IntTestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Int) Filter(f IntTestFunc) Int {
	r := []int{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Int8 []int8

func (s Int8) Equal(v Int8) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type Int8MapFunc func(current int8, index int, array []int8) int8

func (s Int8) Map(f Int8MapFunc) Int8 {
	r := make([]int8, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type Int8ReduceFunc func(previousValue int8, currentValue int8, currentIndex int, array []int8) int8

func (s Int8) Reduce(f Int8ReduceFunc, initialValue int8) int8 {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type Int8TestFunc func(current int8, index int, array []int8) bool

func (s Int8) Some(f Int8TestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Int8) Every(f Int8TestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Int8) Filter(f Int8TestFunc) Int8 {
	r := []int8{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Int16 []int16

func (s Int16) Equal(v Int16) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type Int16MapFunc func(current int16, index int, array []int16) int16

func (s Int16) Map(f Int16MapFunc) Int16 {
	r := make([]int16, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type Int16ReduceFunc func(previousValue int16, currentValue int16, currentIndex int, array []int16) int16

func (s Int16) Reduce(f Int16ReduceFunc, initialValue int16) int16 {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type Int16TestFunc func(current int16, index int, array []int16) bool

func (s Int16) Some(f Int16TestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Int16) Every(f Int16TestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Int16) Filter(f Int16TestFunc) Int16 {
	r := []int16{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Int32 []int32

func (s Int32) Equal(v Int32) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type Int32MapFunc func(current int32, index int, array []int32) int32

func (s Int32) Map(f Int32MapFunc) Int32 {
	r := make([]int32, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type Int32ReduceFunc func(previousValue int32, currentValue int32, currentIndex int, array []int32) int32

func (s Int32) Reduce(f Int32ReduceFunc, initialValue int32) int32 {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type Int32TestFunc func(current int32, index int, array []int32) bool

func (s Int32) Some(f Int32TestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Int32) Every(f Int32TestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Int32) Filter(f Int32TestFunc) Int32 {
	r := []int32{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Int64 []int64

func (s Int64) Equal(v Int64) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type Int64MapFunc func(current int64, index int, array []int64) int64

func (s Int64) Map(f Int64MapFunc) Int64 {
	r := make([]int64, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type Int64ReduceFunc func(previousValue int64, currentValue int64, currentIndex int, array []int64) int64

func (s Int64) Reduce(f Int64ReduceFunc, initialValue int64) int64 {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type Int64TestFunc func(current int64, index int, array []int64) bool

func (s Int64) Some(f Int64TestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Int64) Every(f Int64TestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Int64) Filter(f Int64TestFunc) Int64 {
	r := []int64{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Uint []uint

func (s Uint) Equal(v Uint) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type UintMapFunc func(current uint, index int, array []uint) uint

func (s Uint) Map(f UintMapFunc) Uint {
	r := make([]uint, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type UintReduceFunc func(previousValue uint, currentValue uint, currentIndex int, array []uint) uint

func (s Uint) Reduce(f UintReduceFunc, initialValue uint) uint {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type UintTestFunc func(current uint, index int, array []uint) bool

func (s Uint) Some(f UintTestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Uint) Every(f UintTestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Uint) Filter(f UintTestFunc) Uint {
	r := []uint{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Uint8 []uint8

func (s Uint8) Equal(v Uint8) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type Uint8MapFunc func(current uint8, index int, array []uint8) uint8

func (s Uint8) Map(f Uint8MapFunc) Uint8 {
	r := make([]uint8, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type Uint8ReduceFunc func(previousValue uint8, currentValue uint8, currentIndex int, array []uint8) uint8

func (s Uint8) Reduce(f Uint8ReduceFunc, initialValue uint8) uint8 {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type Uint8TestFunc func(current uint8, index int, array []uint8) bool

func (s Uint8) Some(f Uint8TestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Uint8) Every(f Uint8TestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Uint8) Filter(f Uint8TestFunc) Uint8 {
	r := []uint8{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Uint16 []uint16

func (s Uint16) Equal(v Uint16) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type Uint16MapFunc func(current uint16, index int, array []uint16) uint16

func (s Uint16) Map(f Uint16MapFunc) Uint16 {
	r := make([]uint16, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type Uint16ReduceFunc func(previousValue uint16, currentValue uint16, currentIndex int, array []uint16) uint16

func (s Uint16) Reduce(f Uint16ReduceFunc, initialValue uint16) uint16 {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type Uint16TestFunc func(current uint16, index int, array []uint16) bool

func (s Uint16) Some(f Uint16TestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Uint16) Every(f Uint16TestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Uint16) Filter(f Uint16TestFunc) Uint16 {
	r := []uint16{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Uint32 []uint32

func (s Uint32) Equal(v Uint32) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type Uint32MapFunc func(current uint32, index int, array []uint32) uint32

func (s Uint32) Map(f Uint32MapFunc) Uint32 {
	r := make([]uint32, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type Uint32ReduceFunc func(previousValue uint32, currentValue uint32, currentIndex int, array []uint32) uint32

func (s Uint32) Reduce(f Uint32ReduceFunc, initialValue uint32) uint32 {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type Uint32TestFunc func(current uint32, index int, array []uint32) bool

func (s Uint32) Some(f Uint32TestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Uint32) Every(f Uint32TestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Uint32) Filter(f Uint32TestFunc) Uint32 {
	r := []uint32{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Uint64 []uint64

func (s Uint64) Equal(v Uint64) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type Uint64MapFunc func(current uint64, index int, array []uint64) uint64

func (s Uint64) Map(f Uint64MapFunc) Uint64 {
	r := make([]uint64, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type Uint64ReduceFunc func(previousValue uint64, currentValue uint64, currentIndex int, array []uint64) uint64

func (s Uint64) Reduce(f Uint64ReduceFunc, initialValue uint64) uint64 {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type Uint64TestFunc func(current uint64, index int, array []uint64) bool

func (s Uint64) Some(f Uint64TestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Uint64) Every(f Uint64TestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Uint64) Filter(f Uint64TestFunc) Uint64 {
	r := []uint64{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Rune []rune

func (s Rune) Equal(v Rune) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type RuneMapFunc func(current rune, index int, array []rune) rune

func (s Rune) Map(f RuneMapFunc) Rune {
	r := make([]rune, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type RuneReduceFunc func(previousValue rune, currentValue rune, currentIndex int, array []rune) rune

func (s Rune) Reduce(f RuneReduceFunc, initialValue rune) rune {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type RuneTestFunc func(current rune, index int, array []rune) bool

func (s Rune) Some(f RuneTestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Rune) Every(f RuneTestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Rune) Filter(f RuneTestFunc) Rune {
	r := []rune{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}

type Byte []byte

func (s Byte) Equal(v Byte) bool {
	if s == nil {
		if v == nil {
			return true
		}
		return false
	} else if v == nil {
		return false
	} else if len(s) != len(v) {
		return false
	} else {
		for i, e := range s {
			if e != v[i] {
				return false
			}
		}
		return true
	}
}

type ByteMapFunc func(current byte, index int, array []byte) byte

func (s Byte) Map(f ByteMapFunc) Byte {
	r := make([]byte, len(s))
	for i, e := range s {
		r[i] = f(e, i, s)
	}
	return r
}

type ByteReduceFunc func(previousValue byte, currentValue byte, currentIndex int, array []byte) byte

func (s Byte) Reduce(f ByteReduceFunc, initialValue byte) byte {
	r := initialValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}

type ByteTestFunc func(current byte, index int, array []byte) bool

func (s Byte) Some(f ByteTestFunc) bool {
	for i, e := range s {
		if f(e, i, s) {
			return true
		}
	}
	return false
}

func (s Byte) Every(f ByteTestFunc) bool {
	for i, e := range s {
		if !f(e, i, s) {
			return false
		}
	}
	return true
}

func (s Byte) Filter(f ByteTestFunc) Byte {
	r := []byte{}
	for i, e := range s {
		if f(e, i, s) {
			r = append(r, e)
		}
	}
	return r
}


