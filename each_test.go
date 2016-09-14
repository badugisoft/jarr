package jarr_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/badugisoft/jarr"
)

func TestEach(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}

	jarr.Each(s, func(e interface{}) {
		p := e.(*int)
		*p++
	})

	_true(t, jarr.Equal([]int{2, 3, 4, 5, 6}, s))
}

func ExampleEach() {
	s := []int{1, 2, 3, 4, 5}

	jarr.Each(s, func(e interface{}) {
		p := e.(*int)
		*p++
	})

	fmt.Println(s)
	// output:
	// [2 3 4 5 6]
}

func TestEachIndex(t *testing.T) {
	s := []int{5, 6, 7, 8}

	jarr.EachIndex(s, func(i int, e interface{}) {
		p := e.(*int)
		*p -= i
	})

	_true(t, jarr.Equal([]int{5, 5, 5, 5}, s))
}

func ExampleEachIndex() {
	s := []int{5, 6, 7, 8}

	jarr.EachIndex(s, func(i int, e interface{}) {
		p := e.(*int)
		*p -= i
	})

	fmt.Println(s)
	// output:
	// [5 5 5 5]
}

func TestEachString(t *testing.T) {
	s := []string{"one", "two", "three"}

	jarr.EachString(s, func(e *string) {
		*e = strings.ToUpper((*e)[0:1]) + (*e)[1:]
	})

	_true(t, jarr.Equal([]string{"One", "Two", "Three"}, s))
}

func ExampleEachString() {
	s := []string{"one", "two", "three"}

	jarr.EachString(s, func(e *string) {
		*e = strings.ToUpper((*e)[0:1]) + (*e)[1:]
	})

	fmt.Println(s)
	// output:
	// [One Two Three]
}

func TestEachStringIndex(t *testing.T) {
	s := []string{"one", "two", "three"}

	jarr.EachStringIndex(s, func(i int, e *string) {
		*e = fmt.Sprintf("%v-%v", *e, i+1)
	})

	_true(t, jarr.Equal([]string{"one-1", "two-2", "three-3"}, s))
}

func ExampleEachStringIndex() {
	s := []string{"one", "two", "three"}

	jarr.EachStringIndex(s, func(i int, e *string) {
		*e = fmt.Sprintf("%v-%v", *e, i+1)
	})

	fmt.Println(s)
	// output:
	// [one-1 two-2 three-3]
}
