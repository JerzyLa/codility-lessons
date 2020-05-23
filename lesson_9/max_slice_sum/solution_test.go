package max_slice_sum

import (
	"reflect"
	"testing"
)

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int) int {
	// if all elements are negative return max value
	maxElement := A[0]
	for _, el := range A[1:] {
		maxElement = Max(maxElement, el)
	}

	if maxElement <= 0 {
		return maxElement
	}

	// if any element is positive, find max slice
	maxSlice := 0
	result := 0
	for i := range A {
		maxSlice = Max(0, A[i]+maxSlice)
		result = Max(result, maxSlice)
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A        []int
		Expected int
	}{
		{[]int{3, 2, -6, 4, 0}, 5},
		{[]int{-10}, -10},
	}

	for _, tc := range testCases {
		result := Solution(tc.A)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
