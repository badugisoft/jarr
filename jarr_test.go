package jarr_test

import (
	"fmt"
	"testing"

	"github.com/badugisoft/jarr"
)

func TestSliceEqual(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4, 6}

	_success(t, "s1 == s2", jarr.Equal(s1, s2))
	_fail(t, "s1 == s3", jarr.Equal(s1, s3))
	_fail(t, "s1 == nil", jarr.Equal(s1, nil))
	_success(t, "nil == nil", jarr.Equal(nil, nil))
}

func ExampleSliceEqual() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	fmt.Println(jarr.Equal(s1, s2))
	// output:
	// true
}

func TestEach(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{2, 3, 4, 5, 6}
	jarr.Each(s1, func(e interface{}) {
		p := e.(*int)
		*p++
	})

	_success(t, "s1 == s2", jarr.Equal(s1, s2))
}

func ExampleEach() {
	s1 := []int{1, 2, 3, 4, 5}

	jarr.Each(s1, func(e interface{}) {
		p := e.(*int)
		*p += 3
	})

	fmt.Println(s1)
	// output:
	// [4 5 6 7 8]
}

func TestEachIndex(t *testing.T) {
	s1 := []int{5, 6, 7, 8}
	s2 := []int{5, 5, 5, 5}
	jarr.EachIndex(s1, func(i int, e interface{}) {
		p := e.(*int)
		*p -= i
	})

	_success(t, "s1 == s2", jarr.Equal(s1, s2))
}

func ExampleEachIndex() {
	s1 := []int{5, 6, 7}

	jarr.EachIndex(s1, func(i int, e interface{}) {
		p := e.(*int)
		*p += i * i
	})

	fmt.Println(s1)
	// output:
	// [5 7 11]
}
