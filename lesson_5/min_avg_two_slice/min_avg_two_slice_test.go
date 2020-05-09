package min_avg_two_slice

import (
	"math"
	"reflect"
	"testing"
)

func Solution(A []int) int {
	// write your code in Go 1.4
	min_idx := 0
	min_val := math.MaxFloat64
	for i := range A[:len(A)-1] {
		val := float64(A[i] + A[i+1]) / 2.0
		if val < min_val {
			min_idx = i
			min_val = val
		}

		if i+2 < len(A) {
			val = float64(A[i] + A[i+1] + A[i+2]) / 3.0
			if val < min_val {
				min_idx = i
				min_val = val
			}
		}
	}

	return min_idx
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A              []int
		ExpectedResult int
	}{
		{[]int{4, 2, 2, 5, 1, 5, 8}, 1},
		{[]int{2,3}, 0},
		{[]int{1000, 1000}, 0},
		{[]int{1000,1,1}, 1},
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
