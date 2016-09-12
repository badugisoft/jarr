# jarr
Golang implementation of Javascript array functions like map, reduce, ...


[![GoDoc](https://godoc.org/gopkg.in/badugisoft/jarr.v1?status.svg)](https://godoc.org/gopkg.in/badugisoft/jarr.v1)
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
func ExampleSliceEqual() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	fmt.Println(jarr.Equal(s1, s2))
	// output:
	// true
}
```
