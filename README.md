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
```
