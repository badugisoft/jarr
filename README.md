# jarr
Golang implementation of Javascript array functions like map, reduce, ...


[![GoDoc](https://godoc.org/gopkg.in/badugisoft/jarr.v2?status.svg)](https://godoc.org/gopkg.in/badugisoft/jarr.v2)
[![Build Status](https://travis-ci.org/badugisoft/jarr.svg?branch=v2)](https://travis-ci.org/badugisoft/jarr)
[![Build Status](https://drone.io/github.com/badugisoft/jarr/status.png)](https://drone.io/github.com/badugisoft/jarr/latest)

## Getting started

To get the package, execute:
```bash
go get gopkg.in/badugisoft/jarr.v2
```

To import this package, add the following line to your code:
```go
import "gopkg.in/badugisoft/jarr.v2"
```

## Examples

```go
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
```

```go
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
```

```go
func ExampleEachString() {
	s := []string{"one", "two", "three"}

	jarr.EachString(s, func(e *string) {
		*e = strings.ToUpper((*e)[0:1]) + (*e)[1:]
	})

	fmt.Println(s)
	// output:
	// [One Two Three]
}
```

```go
func ExampleEachStringIndex() {
	s := []string{"one", "two", "three"}

	jarr.EachStringIndex(s, func(i int, e *string) {
		*e = fmt.Sprintf("%v-%v", *e, i+1)
	})

	fmt.Println(s)
	// output:
	// [one-1 two-2 three-3]
}
```

```go
func ExampleFilter() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(jarr.Filter(s, func(e interface{}) bool {
		return e.(int)%2 == 1
	}))
	// output:
	// [1 3 5 7 9]
}
```

```go
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
```

```go
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
```

```go
func ExampleCond() {
	fmt.Println(jarr.Cond(true, 1, 2))
	fmt.Println(jarr.Cond(false, 1, 2))
	// output:
	// 1
	// 2
}
```

```go
func ExamplePrefix() {
	s := []string{"one", "two", "three"}

	fmt.Println(jarr.Prefix(s, "tweenty "))
	// output:
	// [tweenty one tweenty two tweenty three]
}
```

```go
func ExampleSuffix() {
	s := []string{"one", "two", "three"}

	fmt.Println(jarr.Suffix(s, " boy(s)"))
	// output:
	// [one boy(s) two boy(s) three boy(s)]
}
```

```go
func ExampleSuffixPlural() {
	s := []string{"one", "two", "three"}

	fmt.Println(jarr.SuffixPlural(s, " boy", " boys"))
	// output:
	// [one boy two boys three boys]
}
```

```go
func ExampleJoin() {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	fmt.Println(jarr.Join(s1, ","))
	fmt.Println(jarr.Join(s2, ","))
	// output:
	// one,two,three
	// 1,3,5,7
}
```

```go
func ExampleJoinBrace() {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	fmt.Println(jarr.JoinBrace(s1, "[", ",", "]"))
	fmt.Println(jarr.JoinBrace(s2, "<", ",", ">"))
	// output:
	// [one,two,three]
	// <1,3,5,7>
}
```

```go
func ExampleJoinf() {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 3, 5, 7}

	fmt.Println(jarr.Joinf(s1, ",", "%#v"))
	fmt.Println(jarr.Joinf(s2, ",", "%02d"))
	// output:
	// "one","two","three"
	// 01,03,05,07
}
```

```go
func ExampleReverse() {
	s1 := []string{"one", "two", "three"}
	s2 := []int{1, 2, 3}

	fmt.Println(jarr.Reverse(s1))
	fmt.Println(jarr.Reverse(s2))
	// output:
	// [three two one]
	// [3 2 1]
}
```

```go
func ExampleReverseString() {
	s := []string{"one", "two", "three"}

	fmt.Println(jarr.ReverseString(s))
	// output:
	// [three two one]
}
```

```go
func ExampleReverseInt() {
	s := []int{1, 2, 3}

	fmt.Println(jarr.ReverseInt(s))
	// output:
	// [3 2 1]
}
```

```go
func ExampleMap() {
	s1 := []int{5, 6, 7}

	s2 := jarr.Map(s1, func(e interface{}) interface{} {
		return e.(int) + 5
	})

	fmt.Println(s2.([]int))
	// output:
	// [10 11 12]
}
```
