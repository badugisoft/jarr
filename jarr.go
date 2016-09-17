package jarr

import (
	"fmt"
	"reflect"
	"strings"
)

func Equal(s1, s2 interface{}) bool {
	return reflect.DeepEqual(s1, s2)
}

func NotEqual(s1, s2 interface{}) bool {
	return !reflect.DeepEqual(s1, s2)
}

func Cond(c bool, t, f interface{}) interface{} {
	if c {
		return t
	} else {
		return f
	}
}

func Prefix(slice []string, prefix string) []string {
	return Format(slice, Escapef(prefix)+"%s")
}

func Suffix(slice []string, suffix string) []string {
	return Format(slice, "%s"+Escapef(suffix))
}

func SuffixPlural(slice []string, suffix_1, suffix_n string) []string {
	return MapStringIndex(slice, func(i int, e string) string {
		return e + CondString(i == 0, suffix_1, suffix_n)
	})
}

func Escapef(str string) string {
	return strings.Replace(str, "%", "%%", -1)
}

func Format(slice interface{}, format string) []string {
	return MapToString(slice, func(e interface{}) string {
		return fmt.Sprintf(format, e)
	})
}

func Join(slice interface{}, delim string) string {
	return strings.Join(
		MapToString(slice, func(e interface{}) string {
			return fmt.Sprintf("%v", e)
		}), delim)
}

func JoinBrace(slice interface{}, open, delim, close string) string {
	return open + strings.Join(
		MapToString(slice, func(e interface{}) string {
			return fmt.Sprintf("%v", e)
		}), delim) + close
}

func Joinf(slice interface{}, delim, format string) string {
	return strings.Join(
		MapToString(slice, func(e interface{}) string {
			return fmt.Sprintf(format, e)
		}), delim)
}

func JoinfBrace(slice interface{}, open, delim, close, format string) string {
	return open + strings.Join(
		MapToString(slice, func(e interface{}) string {
			return fmt.Sprintf(format, e)
		}), delim) + close
}

func Reverse(slice interface{}) interface{} {
	v := reflect.ValueOf(slice)
	l := v.Len()
	ret := reflect.MakeSlice(reflect.TypeOf(slice), l, l)
	for i := 0; i < l; i++ {
		ret.Index(l - i - 1).Set(v.Index(i))
	}
	return ret.Interface()
}
