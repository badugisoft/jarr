package jarr_test

import (
	"testing"

	"github.com/badugisoft/jarr"
)

func TestJoin(t *testing.T) {
	strs := []string{"apple", "orange", "grape"}
	ints := []int{1, 3, 5, 7, 9}

	_success(t, "String.Join", jarr.String(strs).Join("|") == "apple|orange|grape")
	_success(t, "String.JoinBrace", jarr.String(strs).JoinBrace("(", ",", ")") == "(apple,orange,grape)")
	_success(t, "Int.Join", jarr.Int(ints).Join(", ") == "1, 3, 5, 7, 9")
	_success(t, "Int.JoinBrace", jarr.Int(ints).JoinBrace("[", ",", "]") == "[1,3,5,7,9]")
}
