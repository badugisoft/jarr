package jarr_test

import (
	"fmt"
	"strings"

	"github.com/badugisoft/jarr"
)

func ExampleString_Equal() {
	str1 := []string{"orange", "apple", "banana"}
	str2 := []string{"orange", "apple", "banana"}
	str3 := []string{"orange", "apple", "grape"}

	fmt.Println(jarr.String(str1).Equal(str2))
	fmt.Println(jarr.String(str1).Equal(str3))

	// Output:
	// true
	// false
}

func ExampleString_Map() {
	strs := []string{"one", "two", "three", "four"}
	arr := jarr.String(strs).Map(func(e string) string {
		return strings.ToUpper(e[:1]) + e[1:]
	})

	fmt.Println(arr)

	// Output:
	// [One Two Three Four]
}

func ExampleInt_Filter() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr := jarr.Int(nums).Filter(func(e int) bool {
		return (e % 2) == 1
	})

	fmt.Println(arr)

	// Output:
	// [1 3 5 7 9]
}
