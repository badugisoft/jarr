package jarr

import "testing"

func TestJoin(t *testing.T) {
	strs := []string{"apple", "orange", "grape"}
	ints := []int{1, 3, 5, 7, 9}

	_success(t, "String.Join", String(strs).Join("|") == "apple|orange|grape")
	_success(t, "String.JoinBrace", String(strs).JoinBrace("(", ",", ")") == "(apple,orange,grape)")
	_success(t, "Int.Join", Int(ints).Join(", ") == "1, 3, 5, 7, 9")
	_success(t, "Int.JoinBrace", Int(ints).JoinBrace("[", ",", "]") == "[1,3,5,7,9]")
}
