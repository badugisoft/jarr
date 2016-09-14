package jarr_test

import (
	"regexp"
	"runtime/debug"
	"strings"
	"testing"
)

var reg = regexp.MustCompile("([\\w_]+.go:\\d+)")

func _true(t *testing.T, expr bool) {
	if !expr {
		t.Error(reg.FindString(strings.Split(string(debug.Stack()), "\n")[6]))
	}
}

func _false(t *testing.T, expr bool) {
	if expr {
		t.Error(reg.FindString(strings.Split(string(debug.Stack()), "\n")[6]))
	}
}
