package jarr_test

import (
	"fmt"
	"testing"

	"github.com/badugisoft/jarr"
)

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	_true(t, jarr.Equal(
		[]int{1, 3, 5, 7, 9},
		jarr.Filter(s, func(e interface{}) bool {
			return e.(int)%2 == 1
		})))
}

func ExampleFilter() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(jarr.Filter(s, func(e interface{}) bool {
		return e.(int)%2 == 1
	}))
	// output:
	// [1 3 5 7 9]
}
