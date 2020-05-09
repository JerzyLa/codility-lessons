package max_product_of_three

import (
	"reflect"
	"sort"
	"testing"
)

func Solution(A []int) int {
	sort.Ints(A)
	last := len(A) - 1
	a := A[0] * A[1] * A[last]
	b := A[last-1] * A[last-2] * A[last]

	if a > b {
		return a
	}

	return b
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A              []int
		ExpectedResult int
	}{
		{[]int{-3, 1, 2, -2, 5, 6}, 60},
		{[]int{-3, -3, 0}, 0},
		{[]int{-3, -2, -1, 0, 1}, 6},
	}

	for _, tc := range testCases {
		result := Solution(tc.A)
		if reflect.DeepEqual(result, tc.ExpectedResult) {
			t.Log("PASS")
		} else {
			t.Errorf("FAIL %v, expectes %v", result, tc.ExpectedResult)
		}
	}
}
