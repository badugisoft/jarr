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

type StringMapFunc func(curr string) string

func (s String) Map(f StringMapFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type StringMapIndexFunc func(curr string, index int) string

func (s String) MapIndex(f StringMapIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type StringReduceFunc func(prev string, curr string) string

func (s String) Reduce(f StringReduceFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type StringReduceIndexFunc func(prev string, curr string, index int) string

func (s String) ReduceIndex(f StringReduceIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type StringTestFunc func(current string) bool
type StringTestIndexFunc func(current string, index int) bool

func (s String) Some(f StringTestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s String) SomeIndex(f StringTestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s String) Every(f StringTestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s String) EveryIndex(f StringTestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s String) Filter(f StringTestFunc) String {
	r := []string{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s String) FilterIndex(f StringTestIndexFunc) String {
	r := []string{}
	for i, e := range s {
		if f(e, i) {
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

type BoolMapFunc func(curr bool) bool

func (s Bool) Map(f BoolMapFunc) Bool {
	r := make([]bool, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type BoolMapIndexFunc func(curr bool, index int) bool

func (s Bool) MapIndex(f BoolMapIndexFunc) Bool {
	r := make([]bool, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type BoolReduceFunc func(prev bool, curr bool) bool

func (s Bool) Reduce(f BoolReduceFunc, init bool) bool {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type BoolReduceIndexFunc func(prev bool, curr bool, index int) bool

func (s Bool) ReduceIndex(f BoolReduceIndexFunc, init bool) bool {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type BoolTestFunc func(current bool) bool
type BoolTestIndexFunc func(current bool, index int) bool

func (s Bool) Some(f BoolTestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Bool) SomeIndex(f BoolTestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Bool) Every(f BoolTestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Bool) EveryIndex(f BoolTestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Bool) Filter(f BoolTestFunc) Bool {
	r := []bool{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Bool) FilterIndex(f BoolTestIndexFunc) Bool {
	r := []bool{}
	for i, e := range s {
		if f(e, i) {
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

type IntMapFunc func(curr int) int

func (s Int) Map(f IntMapFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type IntMapIndexFunc func(curr int, index int) int

func (s Int) MapIndex(f IntMapIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type IntReduceFunc func(prev int, curr int) int

func (s Int) Reduce(f IntReduceFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type IntReduceIndexFunc func(prev int, curr int, index int) int

func (s Int) ReduceIndex(f IntReduceIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type IntTestFunc func(current int) bool
type IntTestIndexFunc func(current int, index int) bool

func (s Int) Some(f IntTestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Int) SomeIndex(f IntTestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Int) Every(f IntTestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Int) EveryIndex(f IntTestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Int) Filter(f IntTestFunc) Int {
	r := []int{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Int) FilterIndex(f IntTestIndexFunc) Int {
	r := []int{}
	for i, e := range s {
		if f(e, i) {
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

type Int8MapFunc func(curr int8) int8

func (s Int8) Map(f Int8MapFunc) Int8 {
	r := make([]int8, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int8MapIndexFunc func(curr int8, index int) int8

func (s Int8) MapIndex(f Int8MapIndexFunc) Int8 {
	r := make([]int8, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int8ReduceFunc func(prev int8, curr int8) int8

func (s Int8) Reduce(f Int8ReduceFunc, init int8) int8 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int8ReduceIndexFunc func(prev int8, curr int8, index int) int8

func (s Int8) ReduceIndex(f Int8ReduceIndexFunc, init int8) int8 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int8TestFunc func(current int8) bool
type Int8TestIndexFunc func(current int8, index int) bool

func (s Int8) Some(f Int8TestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Int8) SomeIndex(f Int8TestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Int8) Every(f Int8TestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Int8) EveryIndex(f Int8TestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Int8) Filter(f Int8TestFunc) Int8 {
	r := []int8{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Int8) FilterIndex(f Int8TestIndexFunc) Int8 {
	r := []int8{}
	for i, e := range s {
		if f(e, i) {
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

type Int16MapFunc func(curr int16) int16

func (s Int16) Map(f Int16MapFunc) Int16 {
	r := make([]int16, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int16MapIndexFunc func(curr int16, index int) int16

func (s Int16) MapIndex(f Int16MapIndexFunc) Int16 {
	r := make([]int16, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int16ReduceFunc func(prev int16, curr int16) int16

func (s Int16) Reduce(f Int16ReduceFunc, init int16) int16 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int16ReduceIndexFunc func(prev int16, curr int16, index int) int16

func (s Int16) ReduceIndex(f Int16ReduceIndexFunc, init int16) int16 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int16TestFunc func(current int16) bool
type Int16TestIndexFunc func(current int16, index int) bool

func (s Int16) Some(f Int16TestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Int16) SomeIndex(f Int16TestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Int16) Every(f Int16TestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Int16) EveryIndex(f Int16TestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Int16) Filter(f Int16TestFunc) Int16 {
	r := []int16{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Int16) FilterIndex(f Int16TestIndexFunc) Int16 {
	r := []int16{}
	for i, e := range s {
		if f(e, i) {
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

type Int32MapFunc func(curr int32) int32

func (s Int32) Map(f Int32MapFunc) Int32 {
	r := make([]int32, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int32MapIndexFunc func(curr int32, index int) int32

func (s Int32) MapIndex(f Int32MapIndexFunc) Int32 {
	r := make([]int32, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int32ReduceFunc func(prev int32, curr int32) int32

func (s Int32) Reduce(f Int32ReduceFunc, init int32) int32 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int32ReduceIndexFunc func(prev int32, curr int32, index int) int32

func (s Int32) ReduceIndex(f Int32ReduceIndexFunc, init int32) int32 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int32TestFunc func(current int32) bool
type Int32TestIndexFunc func(current int32, index int) bool

func (s Int32) Some(f Int32TestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Int32) SomeIndex(f Int32TestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Int32) Every(f Int32TestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Int32) EveryIndex(f Int32TestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Int32) Filter(f Int32TestFunc) Int32 {
	r := []int32{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Int32) FilterIndex(f Int32TestIndexFunc) Int32 {
	r := []int32{}
	for i, e := range s {
		if f(e, i) {
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

type Int64MapFunc func(curr int64) int64

func (s Int64) Map(f Int64MapFunc) Int64 {
	r := make([]int64, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int64MapIndexFunc func(curr int64, index int) int64

func (s Int64) MapIndex(f Int64MapIndexFunc) Int64 {
	r := make([]int64, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int64ReduceFunc func(prev int64, curr int64) int64

func (s Int64) Reduce(f Int64ReduceFunc, init int64) int64 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int64ReduceIndexFunc func(prev int64, curr int64, index int) int64

func (s Int64) ReduceIndex(f Int64ReduceIndexFunc, init int64) int64 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int64TestFunc func(current int64) bool
type Int64TestIndexFunc func(current int64, index int) bool

func (s Int64) Some(f Int64TestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Int64) SomeIndex(f Int64TestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Int64) Every(f Int64TestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Int64) EveryIndex(f Int64TestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Int64) Filter(f Int64TestFunc) Int64 {
	r := []int64{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Int64) FilterIndex(f Int64TestIndexFunc) Int64 {
	r := []int64{}
	for i, e := range s {
		if f(e, i) {
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

type UintMapFunc func(curr uint) uint

func (s Uint) Map(f UintMapFunc) Uint {
	r := make([]uint, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type UintMapIndexFunc func(curr uint, index int) uint

func (s Uint) MapIndex(f UintMapIndexFunc) Uint {
	r := make([]uint, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type UintReduceFunc func(prev uint, curr uint) uint

func (s Uint) Reduce(f UintReduceFunc, init uint) uint {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type UintReduceIndexFunc func(prev uint, curr uint, index int) uint

func (s Uint) ReduceIndex(f UintReduceIndexFunc, init uint) uint {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type UintTestFunc func(current uint) bool
type UintTestIndexFunc func(current uint, index int) bool

func (s Uint) Some(f UintTestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Uint) SomeIndex(f UintTestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Uint) Every(f UintTestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Uint) EveryIndex(f UintTestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Uint) Filter(f UintTestFunc) Uint {
	r := []uint{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Uint) FilterIndex(f UintTestIndexFunc) Uint {
	r := []uint{}
	for i, e := range s {
		if f(e, i) {
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

type Uint8MapFunc func(curr uint8) uint8

func (s Uint8) Map(f Uint8MapFunc) Uint8 {
	r := make([]uint8, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint8MapIndexFunc func(curr uint8, index int) uint8

func (s Uint8) MapIndex(f Uint8MapIndexFunc) Uint8 {
	r := make([]uint8, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint8ReduceFunc func(prev uint8, curr uint8) uint8

func (s Uint8) Reduce(f Uint8ReduceFunc, init uint8) uint8 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint8ReduceIndexFunc func(prev uint8, curr uint8, index int) uint8

func (s Uint8) ReduceIndex(f Uint8ReduceIndexFunc, init uint8) uint8 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint8TestFunc func(current uint8) bool
type Uint8TestIndexFunc func(current uint8, index int) bool

func (s Uint8) Some(f Uint8TestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Uint8) SomeIndex(f Uint8TestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Uint8) Every(f Uint8TestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Uint8) EveryIndex(f Uint8TestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Uint8) Filter(f Uint8TestFunc) Uint8 {
	r := []uint8{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Uint8) FilterIndex(f Uint8TestIndexFunc) Uint8 {
	r := []uint8{}
	for i, e := range s {
		if f(e, i) {
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

type Uint16MapFunc func(curr uint16) uint16

func (s Uint16) Map(f Uint16MapFunc) Uint16 {
	r := make([]uint16, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint16MapIndexFunc func(curr uint16, index int) uint16

func (s Uint16) MapIndex(f Uint16MapIndexFunc) Uint16 {
	r := make([]uint16, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint16ReduceFunc func(prev uint16, curr uint16) uint16

func (s Uint16) Reduce(f Uint16ReduceFunc, init uint16) uint16 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint16ReduceIndexFunc func(prev uint16, curr uint16, index int) uint16

func (s Uint16) ReduceIndex(f Uint16ReduceIndexFunc, init uint16) uint16 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint16TestFunc func(current uint16) bool
type Uint16TestIndexFunc func(current uint16, index int) bool

func (s Uint16) Some(f Uint16TestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Uint16) SomeIndex(f Uint16TestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Uint16) Every(f Uint16TestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Uint16) EveryIndex(f Uint16TestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Uint16) Filter(f Uint16TestFunc) Uint16 {
	r := []uint16{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Uint16) FilterIndex(f Uint16TestIndexFunc) Uint16 {
	r := []uint16{}
	for i, e := range s {
		if f(e, i) {
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

type Uint32MapFunc func(curr uint32) uint32

func (s Uint32) Map(f Uint32MapFunc) Uint32 {
	r := make([]uint32, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint32MapIndexFunc func(curr uint32, index int) uint32

func (s Uint32) MapIndex(f Uint32MapIndexFunc) Uint32 {
	r := make([]uint32, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint32ReduceFunc func(prev uint32, curr uint32) uint32

func (s Uint32) Reduce(f Uint32ReduceFunc, init uint32) uint32 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint32ReduceIndexFunc func(prev uint32, curr uint32, index int) uint32

func (s Uint32) ReduceIndex(f Uint32ReduceIndexFunc, init uint32) uint32 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint32TestFunc func(current uint32) bool
type Uint32TestIndexFunc func(current uint32, index int) bool

func (s Uint32) Some(f Uint32TestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Uint32) SomeIndex(f Uint32TestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Uint32) Every(f Uint32TestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Uint32) EveryIndex(f Uint32TestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Uint32) Filter(f Uint32TestFunc) Uint32 {
	r := []uint32{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Uint32) FilterIndex(f Uint32TestIndexFunc) Uint32 {
	r := []uint32{}
	for i, e := range s {
		if f(e, i) {
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

type Uint64MapFunc func(curr uint64) uint64

func (s Uint64) Map(f Uint64MapFunc) Uint64 {
	r := make([]uint64, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint64MapIndexFunc func(curr uint64, index int) uint64

func (s Uint64) MapIndex(f Uint64MapIndexFunc) Uint64 {
	r := make([]uint64, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint64ReduceFunc func(prev uint64, curr uint64) uint64

func (s Uint64) Reduce(f Uint64ReduceFunc, init uint64) uint64 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint64ReduceIndexFunc func(prev uint64, curr uint64, index int) uint64

func (s Uint64) ReduceIndex(f Uint64ReduceIndexFunc, init uint64) uint64 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint64TestFunc func(current uint64) bool
type Uint64TestIndexFunc func(current uint64, index int) bool

func (s Uint64) Some(f Uint64TestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Uint64) SomeIndex(f Uint64TestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Uint64) Every(f Uint64TestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Uint64) EveryIndex(f Uint64TestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Uint64) Filter(f Uint64TestFunc) Uint64 {
	r := []uint64{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Uint64) FilterIndex(f Uint64TestIndexFunc) Uint64 {
	r := []uint64{}
	for i, e := range s {
		if f(e, i) {
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

type RuneMapFunc func(curr rune) rune

func (s Rune) Map(f RuneMapFunc) Rune {
	r := make([]rune, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type RuneMapIndexFunc func(curr rune, index int) rune

func (s Rune) MapIndex(f RuneMapIndexFunc) Rune {
	r := make([]rune, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type RuneReduceFunc func(prev rune, curr rune) rune

func (s Rune) Reduce(f RuneReduceFunc, init rune) rune {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type RuneReduceIndexFunc func(prev rune, curr rune, index int) rune

func (s Rune) ReduceIndex(f RuneReduceIndexFunc, init rune) rune {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type RuneTestFunc func(current rune) bool
type RuneTestIndexFunc func(current rune, index int) bool

func (s Rune) Some(f RuneTestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Rune) SomeIndex(f RuneTestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Rune) Every(f RuneTestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Rune) EveryIndex(f RuneTestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Rune) Filter(f RuneTestFunc) Rune {
	r := []rune{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Rune) FilterIndex(f RuneTestIndexFunc) Rune {
	r := []rune{}
	for i, e := range s {
		if f(e, i) {
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

type ByteMapFunc func(curr byte) byte

func (s Byte) Map(f ByteMapFunc) Byte {
	r := make([]byte, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type ByteMapIndexFunc func(curr byte, index int) byte

func (s Byte) MapIndex(f ByteMapIndexFunc) Byte {
	r := make([]byte, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type ByteReduceFunc func(prev byte, curr byte) byte

func (s Byte) Reduce(f ByteReduceFunc, init byte) byte {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type ByteReduceIndexFunc func(prev byte, curr byte, index int) byte

func (s Byte) ReduceIndex(f ByteReduceIndexFunc, init byte) byte {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type ByteTestFunc func(current byte) bool
type ByteTestIndexFunc func(current byte, index int) bool

func (s Byte) Some(f ByteTestFunc) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

func (s Byte) SomeIndex(f ByteTestIndexFunc) bool {
	for i, e := range s {
		if f(e, i) {
			return true
		}
	}
	return false
}

func (s Byte) Every(f ByteTestFunc) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func (s Byte) EveryIndex(f ByteTestIndexFunc) bool {
	for i, e := range s {
		if !f(e, i) {
			return false
		}
	}
	return true
}

func (s Byte) Filter(f ByteTestFunc) Byte {
	r := []byte{}
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func (s Byte) FilterIndex(f ByteTestIndexFunc) Byte {
	r := []byte{}
	for i, e := range s {
		if f(e, i) {
			r = append(r, e)
		}
	}
	return r
}


