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
	for i, end := 0, len(slice); i < end; i++ {
		ret[i] = prefix + slice[i]
	}
	return ret
}

func Suffix(slice []string, suffix string) []string {
	ret := make([]string, len(slice))
	for i, end := 0, len(slice); i < end; i++ {
		ret[i] = slice[i] + suffix
	}
	return ret
}

func SuffixPlural(slice []string, suffix_1, suffix_n string) []string {
	ret := make([]string, len(slice))
	for i, end := 0, len(slice); i < end; i++ {
		if i == 0 {
			ret[i] = slice[i] + suffix_1
		} else {
			ret[i] = slice[i] + suffix_n
		}
	}
	return ret
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
