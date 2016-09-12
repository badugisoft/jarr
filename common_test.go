package jarr_test

import "testing"

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
