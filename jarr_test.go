package jarr_test

import (
	"fmt"
	"testing"

	"github.com/badugisoft/jarr"
)

func TestEqual(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4, 6}

	_true(t, jarr.Equal(s1, s2))
	_false(t, jarr.Equal(s1, s3))
	_false(t, jarr.Equal(s1, nil))
	_true(t, jarr.Equal(nil, nil))
}

func ExampleEqual() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4, 6}

	fmt.Println(jarr.Equal(s1, s2))
	fmt.Println(jarr.Equal(s1, s3))
	fmt.Println(jarr.Equal(s1, nil))
	fmt.Println(jarr.Equal(nil, nil))
	// output:
	// true
	// false
	// false
	// true
}

func TestNotEqual(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4, 6}

	_false(t, jarr.NotEqual(s1, s2))
	_true(t, jarr.NotEqual(s1, s3))
	_true(t, jarr.NotEqual(s1, nil))
	_false(t, jarr.NotEqual(nil, nil))
}

func ExampleNotEqual() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4, 6}

	fmt.Println(jarr.NotEqual(s1, s2))
	fmt.Println(jarr.NotEqual(s1, s3))
	fmt.Println(jarr.NotEqual(s1, nil))
	fmt.Println(jarr.NotEqual(nil, nil))
	// output:
	// false
	// true
	// true
	// false
}

func TestCond(t *testing.T) {
	_true(t, jarr.Cond(true, 1, 2) == 1)
	_true(t, jarr.Cond(false, 1, 2) == 2)
}

func ExampleCond() {
	fmt.Println(jarr.Cond(true, 1, 2))
	fmt.Println(jarr.Cond(false, 1, 2))
	// output:
	// 1
	// 2
}

func TestPrefix(t *testing.T) {
	s := []string{"one", "two", "three"}

	_true(t, jarr.Equal(
		[]string{"tweenty one", "tweenty two", "tweenty three"},
		jarr.Prefix(s, "tweenty ")))
}

func ExamplePrefix() {
	s := []string{"one", "two", "three"}

	fmt.Println(jarr.Prefix(s, "tweenty "))
	// output:
	// [tweenty one tweenty two tweenty three]
}

func TestSuffix(t *testing.T) {
	s := []string{"one", "two", "three"}

	_true(t, jarr.Equal(
		[]string{"one boy(s)", "two boy(s)", "three boy(s)"},
		jarr.Suffix(s, " boy(s)")))

}

func ExampleSuffix() {
	s := []string{"one", "two", "three"}

	fmt.Println(jarr.Suffix(s, " boy(s)"))
	// output:
	// [one boy(s) two boy(s) three boy(s)]
}

func TestSuffixPlural(t *testing.T) {
	s := []string{"one", "two", "three"}

	_true(t, jarr.Equal(
		[]string{"one boy", "two boys", "three boys"},
		jarr.SuffixPlural(s, " boy", " boys")))
}

func ExampleSuffixPlural() {
	s := []string{"one", "two", "three"}

	fmt.Println(jarr.SuffixPlural(s, " boy", " boys"))
	// output:
	// [one boy two boys three boys]
}

func TestJoin(t *testing.T) {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	_true(t, "one,two,three" == jarr.Join(s1, ","))
	_true(t, "1,3,5,7" == jarr.Join(s2, ","))
}

func ExampleJoin() {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	fmt.Println(jarr.Join(s1, ","))
	fmt.Println(jarr.Join(s2, ","))
	// output:
	// one,two,three
	// 1,3,5,7
}

func TestJoinBrace(t *testing.T) {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	_true(t, "[one,two,three]" == jarr.JoinBrace(s1, "[", ",", "]"))
	_true(t, "<1,3,5,7>" == jarr.JoinBrace(s2, "<", ",", ">"))
}

func ExampleJoinBrace() {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	fmt.Println(jarr.JoinBrace(s1, "[", ",", "]"))
	fmt.Println(jarr.JoinBrace(s2, "<", ",", ">"))
	// output:
	// [one,two,three]
	// <1,3,5,7>
}

func TestJoinf(t *testing.T) {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	_true(t, `"one","two","three"` == jarr.Joinf(s1, ",", "%#v"))
	_true(t, "01,03,05,07" == jarr.Joinf(s2, ",", "%02d"))
}

func ExampleJoinf() {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	fmt.Println(jarr.Joinf(s1, ",", "%#v"))
	fmt.Println(jarr.Joinf(s2, ",", "%02d"))
	// output:
	// "one","two","three"
	// 01,03,05,07
}

func TestReverse(t *testing.T) {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 2, 3}

	_true(t, jarr.Equal([]string{"three", "two", "one"}, jarr.Reverse(s1)))
	_true(t, jarr.Equal([]int{3, 2, 1}, jarr.Reverse(s2)))
}

func ExampleReverse() {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 2, 3}

	fmt.Println(jarr.Reverse(s1))
	fmt.Println(jarr.Reverse(s2))
	// output:
	// [three two one]
	// [3 2 1]
}

func TestReverseString(t *testing.T) {
	s := []string{"one", "two", "three"}

	_true(t, jarr.Equal([]string{"three", "two", "one"}, jarr.ReverseString(s)))
}

func ExampleReverseString() {
	s := []string{"one", "two", "three"}

	fmt.Println(jarr.ReverseString(s))
	// output:
	// [three two one]
}

func TestReverseInt(t *testing.T) {
	s := []int{1, 2, 3}

	_true(t, jarr.Equal([]int{3, 2, 1}, jarr.ReverseInt(s)))
}

func ExampleReverseInt() {
	s := []int{1, 2, 3}

	fmt.Println(jarr.ReverseInt(s))
	// output:
	// [3 2 1]
}
