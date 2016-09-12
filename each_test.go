package jarr_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/badugisoft/jarr"
)

func TestEachString(t *testing.T) {
	s1 := []string{"one", "two", "three"}
	s2 := []string{"One", "Two", "Three"}

	jarr.EachString(s1, func(e *string) {
		*e = strings.ToUpper((*e)[0:1]) + (*e)[1:]
	})

	_success(t, "s1 == s2", jarr.Equal(s1, s2))
}

func ExampleEachString() {
	s1 := []string{"one", "two", "three"}

	jarr.EachString(s1, func(e *string) {
		*e = strings.ToUpper((*e)[0:1]) + (*e)[1:]
	})

	fmt.Println(s1)
	// output:
	// [One Two Three]
}

func TestEachStringIndex(t *testing.T) {
	s1 := []string{"one", "two", "three"}
	s2 := []string{"one-1", "two-2", "three-3"}

	jarr.EachStringIndex(s1, func(i int, e *string) {
		*e = fmt.Sprintf("%v-%v", *e, i+1)
	})

	_success(t, "s1 == s2", jarr.Equal(s1, s2))
}

func ExampleEachStringIndex() {
	s1 := []string{"one", "two", "three"}

	jarr.EachStringIndex(s1, func(i int, e *string) {
		*e = fmt.Sprintf("%v-%v", *e, i+1)
	})

	fmt.Println(s1)
	// output:
	// [one-1 two-2 three-3]
}
