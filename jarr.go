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

type StringMapFunc func(cur string) string

func (s String) Map(f StringMapFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type StringMapIndexFunc func(cur string, idx int) string

func (s String) MapIndex(f StringMapIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type StringMapIntFunc func(cur string) int

func (s String) MapInt(f StringMapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type StringMapIntIndexFunc func(cur string, idx int) int

func (s String) MapIntIndex(f StringMapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type StringReduceFunc func(pre string, cur string) string

func (s String) Reduce(f StringReduceFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type StringReduceIndexFunc func(pre string, cur string, idx int) string

func (s String) ReduceIndex(f StringReduceIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type StringReduceIntFunc func(pre int, cur string) int

func (s String) ReduceInt(f StringReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type StringReduceIntIndexFunc func(pre int, cur string, idx int) int

func (s String) ReduceIntIndex(f StringReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type StringTestFunc func(cur string) bool
type StringTestIndexFunc func(cur string, idx int) bool

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

type BoolMapFunc func(cur bool) bool

func (s Bool) Map(f BoolMapFunc) Bool {
	r := make([]bool, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type BoolMapIndexFunc func(cur bool, idx int) bool

func (s Bool) MapIndex(f BoolMapIndexFunc) Bool {
	r := make([]bool, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type BoolMapStringFunc func(cur bool) string

func (s Bool) MapString(f BoolMapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type BoolMapStringIndexFunc func(cur bool, idx int) string

func (s Bool) MapStringIndex(f BoolMapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type BoolMapIntFunc func(cur bool) int

func (s Bool) MapInt(f BoolMapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type BoolMapIntIndexFunc func(cur bool, idx int) int

func (s Bool) MapIntIndex(f BoolMapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type BoolReduceFunc func(pre bool, cur bool) bool

func (s Bool) Reduce(f BoolReduceFunc, init bool) bool {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type BoolReduceIndexFunc func(pre bool, cur bool, idx int) bool

func (s Bool) ReduceIndex(f BoolReduceIndexFunc, init bool) bool {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type BoolReduceStringFunc func(pre string, cur bool) string

func (s Bool) ReduceString(f BoolReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type BoolReduceStringIndexFunc func(pre string, cur bool, idx int) string

func (s Bool) ReduceStringIndex(f BoolReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type BoolReduceIntFunc func(pre int, cur bool) int

func (s Bool) ReduceInt(f BoolReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type BoolReduceIntIndexFunc func(pre int, cur bool, idx int) int

func (s Bool) ReduceIntIndex(f BoolReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type BoolTestFunc func(cur bool) bool
type BoolTestIndexFunc func(cur bool, idx int) bool

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

type IntMapFunc func(cur int) int

func (s Int) Map(f IntMapFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type IntMapIndexFunc func(cur int, idx int) int

func (s Int) MapIndex(f IntMapIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type IntMapStringFunc func(cur int) string

func (s Int) MapString(f IntMapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type IntMapStringIndexFunc func(cur int, idx int) string

func (s Int) MapStringIndex(f IntMapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type IntReduceFunc func(pre int, cur int) int

func (s Int) Reduce(f IntReduceFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type IntReduceIndexFunc func(pre int, cur int, idx int) int

func (s Int) ReduceIndex(f IntReduceIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type IntReduceStringFunc func(pre string, cur int) string

func (s Int) ReduceString(f IntReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type IntReduceStringIndexFunc func(pre string, cur int, idx int) string

func (s Int) ReduceStringIndex(f IntReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type IntTestFunc func(cur int) bool
type IntTestIndexFunc func(cur int, idx int) bool

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

type Int8MapFunc func(cur int8) int8

func (s Int8) Map(f Int8MapFunc) Int8 {
	r := make([]int8, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int8MapIndexFunc func(cur int8, idx int) int8

func (s Int8) MapIndex(f Int8MapIndexFunc) Int8 {
	r := make([]int8, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int8MapStringFunc func(cur int8) string

func (s Int8) MapString(f Int8MapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int8MapStringIndexFunc func(cur int8, idx int) string

func (s Int8) MapStringIndex(f Int8MapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type Int8MapIntFunc func(cur int8) int

func (s Int8) MapInt(f Int8MapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int8MapIntIndexFunc func(cur int8, idx int) int

func (s Int8) MapIntIndex(f Int8MapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int8ReduceFunc func(pre int8, cur int8) int8

func (s Int8) Reduce(f Int8ReduceFunc, init int8) int8 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int8ReduceIndexFunc func(pre int8, cur int8, idx int) int8

func (s Int8) ReduceIndex(f Int8ReduceIndexFunc, init int8) int8 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int8ReduceStringFunc func(pre string, cur int8) string

func (s Int8) ReduceString(f Int8ReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int8ReduceStringIndexFunc func(pre string, cur int8, idx int) string

func (s Int8) ReduceStringIndex(f Int8ReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type Int8ReduceIntFunc func(pre int, cur int8) int

func (s Int8) ReduceInt(f Int8ReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int8ReduceIntIndexFunc func(pre int, cur int8, idx int) int

func (s Int8) ReduceIntIndex(f Int8ReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int8TestFunc func(cur int8) bool
type Int8TestIndexFunc func(cur int8, idx int) bool

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

type Int16MapFunc func(cur int16) int16

func (s Int16) Map(f Int16MapFunc) Int16 {
	r := make([]int16, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int16MapIndexFunc func(cur int16, idx int) int16

func (s Int16) MapIndex(f Int16MapIndexFunc) Int16 {
	r := make([]int16, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int16MapStringFunc func(cur int16) string

func (s Int16) MapString(f Int16MapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int16MapStringIndexFunc func(cur int16, idx int) string

func (s Int16) MapStringIndex(f Int16MapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type Int16MapIntFunc func(cur int16) int

func (s Int16) MapInt(f Int16MapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int16MapIntIndexFunc func(cur int16, idx int) int

func (s Int16) MapIntIndex(f Int16MapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int16ReduceFunc func(pre int16, cur int16) int16

func (s Int16) Reduce(f Int16ReduceFunc, init int16) int16 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int16ReduceIndexFunc func(pre int16, cur int16, idx int) int16

func (s Int16) ReduceIndex(f Int16ReduceIndexFunc, init int16) int16 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int16ReduceStringFunc func(pre string, cur int16) string

func (s Int16) ReduceString(f Int16ReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int16ReduceStringIndexFunc func(pre string, cur int16, idx int) string

func (s Int16) ReduceStringIndex(f Int16ReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type Int16ReduceIntFunc func(pre int, cur int16) int

func (s Int16) ReduceInt(f Int16ReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int16ReduceIntIndexFunc func(pre int, cur int16, idx int) int

func (s Int16) ReduceIntIndex(f Int16ReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int16TestFunc func(cur int16) bool
type Int16TestIndexFunc func(cur int16, idx int) bool

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

type Int32MapFunc func(cur int32) int32

func (s Int32) Map(f Int32MapFunc) Int32 {
	r := make([]int32, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int32MapIndexFunc func(cur int32, idx int) int32

func (s Int32) MapIndex(f Int32MapIndexFunc) Int32 {
	r := make([]int32, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int32MapStringFunc func(cur int32) string

func (s Int32) MapString(f Int32MapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int32MapStringIndexFunc func(cur int32, idx int) string

func (s Int32) MapStringIndex(f Int32MapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type Int32MapIntFunc func(cur int32) int

func (s Int32) MapInt(f Int32MapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int32MapIntIndexFunc func(cur int32, idx int) int

func (s Int32) MapIntIndex(f Int32MapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int32ReduceFunc func(pre int32, cur int32) int32

func (s Int32) Reduce(f Int32ReduceFunc, init int32) int32 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int32ReduceIndexFunc func(pre int32, cur int32, idx int) int32

func (s Int32) ReduceIndex(f Int32ReduceIndexFunc, init int32) int32 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int32ReduceStringFunc func(pre string, cur int32) string

func (s Int32) ReduceString(f Int32ReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int32ReduceStringIndexFunc func(pre string, cur int32, idx int) string

func (s Int32) ReduceStringIndex(f Int32ReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type Int32ReduceIntFunc func(pre int, cur int32) int

func (s Int32) ReduceInt(f Int32ReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int32ReduceIntIndexFunc func(pre int, cur int32, idx int) int

func (s Int32) ReduceIntIndex(f Int32ReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int32TestFunc func(cur int32) bool
type Int32TestIndexFunc func(cur int32, idx int) bool

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

type Int64MapFunc func(cur int64) int64

func (s Int64) Map(f Int64MapFunc) Int64 {
	r := make([]int64, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int64MapIndexFunc func(cur int64, idx int) int64

func (s Int64) MapIndex(f Int64MapIndexFunc) Int64 {
	r := make([]int64, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int64MapStringFunc func(cur int64) string

func (s Int64) MapString(f Int64MapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int64MapStringIndexFunc func(cur int64, idx int) string

func (s Int64) MapStringIndex(f Int64MapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type Int64MapIntFunc func(cur int64) int

func (s Int64) MapInt(f Int64MapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Int64MapIntIndexFunc func(cur int64, idx int) int

func (s Int64) MapIntIndex(f Int64MapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Int64ReduceFunc func(pre int64, cur int64) int64

func (s Int64) Reduce(f Int64ReduceFunc, init int64) int64 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int64ReduceIndexFunc func(pre int64, cur int64, idx int) int64

func (s Int64) ReduceIndex(f Int64ReduceIndexFunc, init int64) int64 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int64ReduceStringFunc func(pre string, cur int64) string

func (s Int64) ReduceString(f Int64ReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int64ReduceStringIndexFunc func(pre string, cur int64, idx int) string

func (s Int64) ReduceStringIndex(f Int64ReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type Int64ReduceIntFunc func(pre int, cur int64) int

func (s Int64) ReduceInt(f Int64ReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Int64ReduceIntIndexFunc func(pre int, cur int64, idx int) int

func (s Int64) ReduceIntIndex(f Int64ReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Int64TestFunc func(cur int64) bool
type Int64TestIndexFunc func(cur int64, idx int) bool

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

type UintMapFunc func(cur uint) uint

func (s Uint) Map(f UintMapFunc) Uint {
	r := make([]uint, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type UintMapIndexFunc func(cur uint, idx int) uint

func (s Uint) MapIndex(f UintMapIndexFunc) Uint {
	r := make([]uint, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type UintMapStringFunc func(cur uint) string

func (s Uint) MapString(f UintMapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type UintMapStringIndexFunc func(cur uint, idx int) string

func (s Uint) MapStringIndex(f UintMapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type UintMapIntFunc func(cur uint) int

func (s Uint) MapInt(f UintMapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type UintMapIntIndexFunc func(cur uint, idx int) int

func (s Uint) MapIntIndex(f UintMapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type UintReduceFunc func(pre uint, cur uint) uint

func (s Uint) Reduce(f UintReduceFunc, init uint) uint {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type UintReduceIndexFunc func(pre uint, cur uint, idx int) uint

func (s Uint) ReduceIndex(f UintReduceIndexFunc, init uint) uint {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type UintReduceStringFunc func(pre string, cur uint) string

func (s Uint) ReduceString(f UintReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type UintReduceStringIndexFunc func(pre string, cur uint, idx int) string

func (s Uint) ReduceStringIndex(f UintReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type UintReduceIntFunc func(pre int, cur uint) int

func (s Uint) ReduceInt(f UintReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type UintReduceIntIndexFunc func(pre int, cur uint, idx int) int

func (s Uint) ReduceIntIndex(f UintReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type UintTestFunc func(cur uint) bool
type UintTestIndexFunc func(cur uint, idx int) bool

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

type Uint8MapFunc func(cur uint8) uint8

func (s Uint8) Map(f Uint8MapFunc) Uint8 {
	r := make([]uint8, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint8MapIndexFunc func(cur uint8, idx int) uint8

func (s Uint8) MapIndex(f Uint8MapIndexFunc) Uint8 {
	r := make([]uint8, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint8MapStringFunc func(cur uint8) string

func (s Uint8) MapString(f Uint8MapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint8MapStringIndexFunc func(cur uint8, idx int) string

func (s Uint8) MapStringIndex(f Uint8MapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type Uint8MapIntFunc func(cur uint8) int

func (s Uint8) MapInt(f Uint8MapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint8MapIntIndexFunc func(cur uint8, idx int) int

func (s Uint8) MapIntIndex(f Uint8MapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint8ReduceFunc func(pre uint8, cur uint8) uint8

func (s Uint8) Reduce(f Uint8ReduceFunc, init uint8) uint8 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint8ReduceIndexFunc func(pre uint8, cur uint8, idx int) uint8

func (s Uint8) ReduceIndex(f Uint8ReduceIndexFunc, init uint8) uint8 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint8ReduceStringFunc func(pre string, cur uint8) string

func (s Uint8) ReduceString(f Uint8ReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint8ReduceStringIndexFunc func(pre string, cur uint8, idx int) string

func (s Uint8) ReduceStringIndex(f Uint8ReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type Uint8ReduceIntFunc func(pre int, cur uint8) int

func (s Uint8) ReduceInt(f Uint8ReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint8ReduceIntIndexFunc func(pre int, cur uint8, idx int) int

func (s Uint8) ReduceIntIndex(f Uint8ReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint8TestFunc func(cur uint8) bool
type Uint8TestIndexFunc func(cur uint8, idx int) bool

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

type Uint16MapFunc func(cur uint16) uint16

func (s Uint16) Map(f Uint16MapFunc) Uint16 {
	r := make([]uint16, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint16MapIndexFunc func(cur uint16, idx int) uint16

func (s Uint16) MapIndex(f Uint16MapIndexFunc) Uint16 {
	r := make([]uint16, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint16MapStringFunc func(cur uint16) string

func (s Uint16) MapString(f Uint16MapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint16MapStringIndexFunc func(cur uint16, idx int) string

func (s Uint16) MapStringIndex(f Uint16MapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type Uint16MapIntFunc func(cur uint16) int

func (s Uint16) MapInt(f Uint16MapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint16MapIntIndexFunc func(cur uint16, idx int) int

func (s Uint16) MapIntIndex(f Uint16MapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint16ReduceFunc func(pre uint16, cur uint16) uint16

func (s Uint16) Reduce(f Uint16ReduceFunc, init uint16) uint16 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint16ReduceIndexFunc func(pre uint16, cur uint16, idx int) uint16

func (s Uint16) ReduceIndex(f Uint16ReduceIndexFunc, init uint16) uint16 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint16ReduceStringFunc func(pre string, cur uint16) string

func (s Uint16) ReduceString(f Uint16ReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint16ReduceStringIndexFunc func(pre string, cur uint16, idx int) string

func (s Uint16) ReduceStringIndex(f Uint16ReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type Uint16ReduceIntFunc func(pre int, cur uint16) int

func (s Uint16) ReduceInt(f Uint16ReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint16ReduceIntIndexFunc func(pre int, cur uint16, idx int) int

func (s Uint16) ReduceIntIndex(f Uint16ReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint16TestFunc func(cur uint16) bool
type Uint16TestIndexFunc func(cur uint16, idx int) bool

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

type Uint32MapFunc func(cur uint32) uint32

func (s Uint32) Map(f Uint32MapFunc) Uint32 {
	r := make([]uint32, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint32MapIndexFunc func(cur uint32, idx int) uint32

func (s Uint32) MapIndex(f Uint32MapIndexFunc) Uint32 {
	r := make([]uint32, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint32MapStringFunc func(cur uint32) string

func (s Uint32) MapString(f Uint32MapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint32MapStringIndexFunc func(cur uint32, idx int) string

func (s Uint32) MapStringIndex(f Uint32MapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type Uint32MapIntFunc func(cur uint32) int

func (s Uint32) MapInt(f Uint32MapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint32MapIntIndexFunc func(cur uint32, idx int) int

func (s Uint32) MapIntIndex(f Uint32MapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint32ReduceFunc func(pre uint32, cur uint32) uint32

func (s Uint32) Reduce(f Uint32ReduceFunc, init uint32) uint32 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint32ReduceIndexFunc func(pre uint32, cur uint32, idx int) uint32

func (s Uint32) ReduceIndex(f Uint32ReduceIndexFunc, init uint32) uint32 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint32ReduceStringFunc func(pre string, cur uint32) string

func (s Uint32) ReduceString(f Uint32ReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint32ReduceStringIndexFunc func(pre string, cur uint32, idx int) string

func (s Uint32) ReduceStringIndex(f Uint32ReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type Uint32ReduceIntFunc func(pre int, cur uint32) int

func (s Uint32) ReduceInt(f Uint32ReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint32ReduceIntIndexFunc func(pre int, cur uint32, idx int) int

func (s Uint32) ReduceIntIndex(f Uint32ReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint32TestFunc func(cur uint32) bool
type Uint32TestIndexFunc func(cur uint32, idx int) bool

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

type Uint64MapFunc func(cur uint64) uint64

func (s Uint64) Map(f Uint64MapFunc) Uint64 {
	r := make([]uint64, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint64MapIndexFunc func(cur uint64, idx int) uint64

func (s Uint64) MapIndex(f Uint64MapIndexFunc) Uint64 {
	r := make([]uint64, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint64MapStringFunc func(cur uint64) string

func (s Uint64) MapString(f Uint64MapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint64MapStringIndexFunc func(cur uint64, idx int) string

func (s Uint64) MapStringIndex(f Uint64MapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type Uint64MapIntFunc func(cur uint64) int

func (s Uint64) MapInt(f Uint64MapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type Uint64MapIntIndexFunc func(cur uint64, idx int) int

func (s Uint64) MapIntIndex(f Uint64MapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type Uint64ReduceFunc func(pre uint64, cur uint64) uint64

func (s Uint64) Reduce(f Uint64ReduceFunc, init uint64) uint64 {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint64ReduceIndexFunc func(pre uint64, cur uint64, idx int) uint64

func (s Uint64) ReduceIndex(f Uint64ReduceIndexFunc, init uint64) uint64 {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint64ReduceStringFunc func(pre string, cur uint64) string

func (s Uint64) ReduceString(f Uint64ReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint64ReduceStringIndexFunc func(pre string, cur uint64, idx int) string

func (s Uint64) ReduceStringIndex(f Uint64ReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type Uint64ReduceIntFunc func(pre int, cur uint64) int

func (s Uint64) ReduceInt(f Uint64ReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type Uint64ReduceIntIndexFunc func(pre int, cur uint64, idx int) int

func (s Uint64) ReduceIntIndex(f Uint64ReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type Uint64TestFunc func(cur uint64) bool
type Uint64TestIndexFunc func(cur uint64, idx int) bool

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

type RuneMapFunc func(cur rune) rune

func (s Rune) Map(f RuneMapFunc) Rune {
	r := make([]rune, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type RuneMapIndexFunc func(cur rune, idx int) rune

func (s Rune) MapIndex(f RuneMapIndexFunc) Rune {
	r := make([]rune, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type RuneMapStringFunc func(cur rune) string

func (s Rune) MapString(f RuneMapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type RuneMapStringIndexFunc func(cur rune, idx int) string

func (s Rune) MapStringIndex(f RuneMapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type RuneMapIntFunc func(cur rune) int

func (s Rune) MapInt(f RuneMapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type RuneMapIntIndexFunc func(cur rune, idx int) int

func (s Rune) MapIntIndex(f RuneMapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type RuneReduceFunc func(pre rune, cur rune) rune

func (s Rune) Reduce(f RuneReduceFunc, init rune) rune {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type RuneReduceIndexFunc func(pre rune, cur rune, idx int) rune

func (s Rune) ReduceIndex(f RuneReduceIndexFunc, init rune) rune {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type RuneReduceStringFunc func(pre string, cur rune) string

func (s Rune) ReduceString(f RuneReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type RuneReduceStringIndexFunc func(pre string, cur rune, idx int) string

func (s Rune) ReduceStringIndex(f RuneReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type RuneReduceIntFunc func(pre int, cur rune) int

func (s Rune) ReduceInt(f RuneReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type RuneReduceIntIndexFunc func(pre int, cur rune, idx int) int

func (s Rune) ReduceIntIndex(f RuneReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type RuneTestFunc func(cur rune) bool
type RuneTestIndexFunc func(cur rune, idx int) bool

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

type ByteMapFunc func(cur byte) byte

func (s Byte) Map(f ByteMapFunc) Byte {
	r := make([]byte, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type ByteMapIndexFunc func(cur byte, idx int) byte

func (s Byte) MapIndex(f ByteMapIndexFunc) Byte {
	r := make([]byte, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type ByteMapStringFunc func(cur byte) string

func (s Byte) MapString(f ByteMapStringFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type ByteMapStringIndexFunc func(cur byte, idx int) string

func (s Byte) MapStringIndex(f ByteMapStringIndexFunc) String {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}


type ByteMapIntFunc func(cur byte) int

func (s Byte) MapInt(f ByteMapIntFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

type ByteMapIntIndexFunc func(cur byte, idx int) int

func (s Byte) MapIntIndex(f ByteMapIntIndexFunc) Int {
	r := make([]int, len(s))
	for i, e := range s {
		r[i] = f(e, i)
	}
	return r
}

type ByteReduceFunc func(pre byte, cur byte) byte

func (s Byte) Reduce(f ByteReduceFunc, init byte) byte {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type ByteReduceIndexFunc func(pre byte, cur byte, idx int) byte

func (s Byte) ReduceIndex(f ByteReduceIndexFunc, init byte) byte {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type ByteReduceStringFunc func(pre string, cur byte) string

func (s Byte) ReduceString(f ByteReduceStringFunc, init string) string {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type ByteReduceStringIndexFunc func(pre string, cur byte, idx int) string

func (s Byte) ReduceStringIndex(f ByteReduceStringIndexFunc, init string) string {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}


type ByteReduceIntFunc func(pre int, cur byte) int

func (s Byte) ReduceInt(f ByteReduceIntFunc, init int) int {
	r := init
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

type ByteReduceIntIndexFunc func(pre int, cur byte, idx int) int

func (s Byte) ReduceIntIndex(f ByteReduceIntIndexFunc, init int) int {
	r := init
	for i, e := range s {
		r = f(r, e, i)
	}
	return r
}

type ByteTestFunc func(cur byte) bool
type ByteTestIndexFunc func(cur byte, idx int) bool

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


