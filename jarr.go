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

func Prefix(slice []string, prefix string) []string {
	ret := make([]string, len(slice))
	for i, e := range slice {
		ret[i] = prefix + e
	}
	return ret
}

func Suffix(slice []string, suffix string) []string {
	ret := make([]string, len(slice))
	for i, e := range slice {
		ret[i] = e + suffix
	}
	return ret
}

func SuffixPlural(slice []string, suffix_1, suffix_n string) []string {
	ret := make([]string, len(slice))
	for i, e := range slice {
		if i == 0 {
			ret[i] = e + suffix_1
		} else {
			ret[i] = e + suffix_n
		}
	}
	return ret
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
