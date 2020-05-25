package count_factors

import (
	"reflect"
	"testing"
)

func Solution(N int) int {
	result := 0

	var i = 1
	for ; i*i < N; i++ {
		if N%i == 0 {
			result += 2
		}
	}

	if i*i == N {
		result++
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{24, 8},
	}

	for _, tc := range testCases {
		result := Solution(tc.N)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
