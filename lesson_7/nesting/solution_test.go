package nesting

import (
	"reflect"
	"testing"
)

func Solution(S string) int {
	if S == "" {
		return 1
	} else if S[0] == ')' {
		return 0
	}

	res := 0
	for _, c := range S {
		if c == '(' {
			res++
		} else {
			res--
		}
		if res < 0 {
			return 0
		}
	}

	if res != 0 {
		return 0
	}
	return 1
}

func TestSolution(t *testing.T) {
	testCases := []struct{
		S string
		Expected int
	} {
		{"(()(())())", 1},
		{"())", 0},
		{"", 1},
		{"()", 1},
		{")(", 0},
		{"())(()", 0},
	}

	for _, tc := range testCases {
		result := Solution(tc.S)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Log("PASS")
		} else {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
