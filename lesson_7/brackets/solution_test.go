package brackets

import (
	"reflect"
	"testing"
)

func Solution(S string) int {
	if S == "" {
		return 1
	}

	var stack []int32
	for _, c := range S {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0{
				return 0
			}
			n := len(stack)-1
			prev := stack[n]
			if (c == '}' && prev != '{') || (c == ')' && prev != '(') || (c == ']' && prev != '[') {
				return 0
			} else {
				stack = stack[:n]
			}
		}
	}

	if len(stack) == 0 {
		return 1
	}
	return 0
}

func TestSolution(t *testing.T) {
	testCases := []struct{
		A string
		Expected int
	} {
		{"{[()()]}", 1},
		{"([)()]", 0},
		{"", 1},
		{"{", 0},
		{"}", 0},
		{"{}", 1},
		{"{]", 0},
	}

	for _, tc := range testCases {
		result := Solution(tc.A)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Log("PASS")
		} else {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
