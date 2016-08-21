package jarr

import (
	"fmt"
	"strings"
	"testing"
)

func _success(t *testing.T, desc string, expr bool) {
	if !expr {
		t.Error(desc)
	}
}

func _fail(t *testing.T, desc string, expr bool) {
	if expr {
		t.Error(desc)
	}
}

func TestStringEqual(t *testing.T) {
	str := []string{"apple", "orange", "grape"}
	ok := []string{"apple", "orange", "grape"}
	fail := [][]string{
		[]string{"apple", "orange", "grapes"},
		[]string{"apple", "orange"},
		[]string{"apple", "orange", "grapes"},
		[]string{"apple", "orange", "grape", "pineapple"},
	}

	_success(t, "nil == nil", String(nil).Equal(nil))
	_fail(t, "nil == str", String(nil).Equal(str))
	_fail(t, "str == nil", String(str).Equal(nil))
	_success(t, "str == ok", String(str).Equal(ok))

	for i, e := range fail {
		_fail(t, fmt.Sprintf("nil == fail[%v]", i), String(str).Equal(e))
	}

}

func TestStringMap(t *testing.T) {
	in := []string{"apple", "orange", "grape"}
	ok := []string{"Apples", "Oranges", "Grapes"}
	res := String(in).Map(func(e string) string {
		return fmt.Sprintf("%v%vs", strings.ToUpper(e[0:1]), e[1:])
	})

	_success(t, "res == ok", String(res).Equal(ok))
}

func TestStringReduce(t *testing.T) {
	in := []string{"apple", "orange", "grape"}
	ok := "start-apple-orange-grape"
	res := String(in).Reduce(func(p string, e string) string {
		return fmt.Sprintf("%v-%v", p, e)
	}, "start")

	_success(t, "res == ok", res == ok)
}

func TestStringSome(t *testing.T) {
	in := []string{"apple", "orange", "grape"}

	_success(t, "len(e) > 5", String(in).Some(func(e string) bool { return len(e) > 5 }))
	_fail(t, "len(e) > 6", String(in).Some(func(e string) bool { return len(e) > 6 }))
	_success(t, "e == \"apple\"", String(in).Some(func(e string) bool { return e == "apple" }))
}

func TestStringEvery(t *testing.T) {
	in := []string{"apple", "orange", "grape"}

	_success(t, "len(e) > 4", String(in).Every(func(e string) bool { return len(e) > 4 }))
	_fail(t, "len(e) > 5", String(in).Every(func(e string) bool { return len(e) > 5 }))
	_fail(t, "e == \"apple\"", String(in).Every(func(e string) bool { return e == "apple" }))
}

func TestStringFilter(t *testing.T) {
	in := []string{"apple", "orange", "grape"}

	_success(t, "len(e) > 5",
		String(in).Filter(func(e string) bool {
			return len(e) > 5
		}).Equal([]string{"orange"}))

	_success(t, "e[0:1] == \"g\"",
		String(in).Filter(func(e string) bool {
			return e[0:1] == "g"
		}).Equal([]string{"grape"}))
}
