package fish

import (
	"reflect"
	"testing"
)

func Solution(A []int, B []int) int {
	var stack []int
	alive := 0

	for i := range A{
		if B[i] == 1 {
			stack = append(stack, A[i])
		} else if len(stack) != 0 {
			n := len(stack)-1
			for ; len(stack) != 0 && stack[n] < A[i] ; n-- {
				stack = stack[:n]
				alive--
			}
			if len(stack) != 0 && stack[n] > A[i] {
				alive--
			}
		}
		alive++
	}

	return alive
}

func TestSolution(t *testing.T) {
	testCases := []struct{
		A []int
		B []int
		Expected int
	} {
		{[]int{4,3,2,1,5}, []int{0,1,0,0,0}, 2},
	}

	for _, tc := range testCases {
		result := Solution(tc.A, tc.B)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Log("PASS")
		} else {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
