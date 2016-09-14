package jarr_test

import (
	"fmt"
	"testing"

	"github.com/badugisoft/jarr"
)

func TestMap(t *testing.T) {
	s1 := []int{2, 1, 3}
	s2 := []int{4, 1, 9}

	s3 := jarr.Map(s1, func(e interface{}) interface{} {
		return e.(int) * e.(int)
	})

	_true(t, jarr.Equal(s3, s2))
}

func ExampleMap() {
	s1 := []int{5, 6, 7}

	s2 := jarr.Map(s1, func(e interface{}) interface{} {
		return e.(int) + 5
	})

	fmt.Println(s2.([]int))
	// output:
	// [10 11 12]
}

func TestMapType(t *testing.T) {
	s1 := []int{2, 1, 3}
	s2 := []int64{4, 1, 9}

	s3 := jarr.MapType(s1, []int64{}, func(e interface{}) interface{} {
		return int64(e.(int) * e.(int))
	})

	_true(t, jarr.Equal(s3, s2))
}
