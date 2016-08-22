package jarr_test

import (
	"fmt"

	"github.com/badugisoft/jarr"
)

func ExampleString_Equal() {
	str1 := []string{"orange", "apple", "banana"}
	str2 := []string{"orange", "apple", "banana"}
	str3 := []string{"orange", "apple", "grape"}

	fmt.Println(jarr.String(str1).Equal(str2))
	fmt.Println(jarr.String(str1).Equal(str3))
}
