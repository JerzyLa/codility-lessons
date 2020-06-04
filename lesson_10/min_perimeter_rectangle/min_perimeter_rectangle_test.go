package min_perimeter_rectangle

import (
	"reflect"
	"testing"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solution(N int) int {
	result := 2 * (N + 1)
	for i := 2; i*i <= N; i++ {
		if N%i == 0 {
			result = Min(result, 2*(i+N/i))
		}
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{30, 22},
	}

	for _, tc := range testCases {
		result := Solution(tc.N)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
