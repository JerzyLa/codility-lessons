package max_double_slice_sum

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
	maxSumStart := make([]int, len(A))
	maxSumEnd := make([]int, len(A))

	for i := 1; i < len(A)-1; i++ {
		maxSumEnd[i] = Max(0, maxSumEnd[i-1]+A[i])
	}
	for i := len(A) - 2; i > 0; i-- {
		maxSumStart[i] = Max(0, maxSumStart[i+1]+A[i])
	}

	result := 0
	for i := 1; i < len(A)-1; i++ {
		temp := maxSumEnd[i-1] + maxSumStart[i+1]
		result = Max(result, temp)
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A        []int
		Expected int
	}{
		{[]int{3, 2, 6, -1, 4, 5, -1, 2}, 17},
	}

	for _, tc := range testCases {
		result := Solution(tc.A)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
