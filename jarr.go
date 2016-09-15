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

func JoinBrace(slice interface{}, start, delim, end string) string {
	return start + strings.Join(
		MapToString(slice, func(e interface{}) string {
			return fmt.Sprintf("%v", e)
		}), delim) + end
}

func Joinf(slice interface{}, delim, format string) string {
	return strings.Join(
		MapToString(slice, func(e interface{}) string {
			return fmt.Sprintf(format, e)
		}), delim)
}

func JoinfBrace(slice interface{}, start, delim, end, format string) string {
	return start + strings.Join(
		MapToString(slice, func(e interface{}) string {
			return fmt.Sprintf(format, e)
		}), delim) + end
}

func Reverse(slice interface{}) interface{} {
	return nil
}

func ReverseString(slice []string) []string {
	l := len(slice)
	for i, end := 0, l/2; i < end; i++ {
		slice[i], slice[l-i-1] = slice[l-i-1], slice[i]
	}
	return slice
}
