package max_profit

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

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	minimum := make([]int, len(A))

	minimum[0] = A[0]
	for i := 1; i < len(A); i++ {
		minimum[i] = Min(minimum[i-1], A[i])
	}

	result := 0
	for i := 1; i < len(A); i++ {
		temp := A[i] - minimum[i-1]
		result = Max(result, temp)
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A        []int
		Expected int
	}{
		{[]int{23171, 21011, 21123, 21366, 21013, 21367}, 356},
	}

	for _, tc := range testCases {
		result := Solution(tc.A)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
