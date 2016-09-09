package jarr

import (
	"strconv"
	"strings"
)

func (s String) Join(delim string) string {
	return strings.Join(s, delim)
}

func (s String) JoinBrace(open, delim, close string) string {
	return open + strings.Join(s, delim) + close
}

func (s Int) Join(delim string) string {
	return s.MapString(func(e int) string {
		return strconv.FormatInt(int64(e), 10)
	}).Join(delim)
}

func (s Int) JoinBrace(open, delim, close string) string {
	return s.MapString(func(e int) string {
		return strconv.FormatInt(int64(e), 10)
	}).JoinBrace(open, delim, close)
}

func (s Int64) Join(delim string) string {
	return s.MapString(func(e int64) string {
		return strconv.FormatInt(e, 10)
	}).Join(delim)
}

func (s Int64) JoinBrace(open, delim, close string) string {
	return s.MapString(func(e int64) string {
		return strconv.FormatInt(e, 10)
	}).JoinBrace(open, delim, close)
}

func (s Uint) Join(delim string) string {
	return s.MapString(func(e uint) string {
		return strconv.FormatUint(uint64(e), 10)
	}).Join(delim)
}

func (s Uint) JoinBrace(open, delim, close string) string {
	return s.MapString(func(e uint) string {
		return strconv.FormatUint(uint64(e), 10)
	}).JoinBrace(open, delim, close)
}

func (s Uint64) Join(delim string) string {
	return s.MapString(func(e uint64) string {
		return strconv.FormatUint(e, 10)
	}).Join(delim)
}

func (s Uint64) JoinBrace(open, delim, close string) string {
	return s.MapString(func(e uint64) string {
		return strconv.FormatUint(e, 10)
	}).JoinBrace(open, delim, close)
}

func (s String) Prefix(str string) String {
	return s.Map(func(e string) string {
		return str + e
	})
}

func (s String) Suffix(str string) String {
	return s.Map(func(e string) string {
		return e + str
	})
}
