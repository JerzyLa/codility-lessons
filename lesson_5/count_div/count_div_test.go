package count_div

import (
	"reflect"
	"testing"
)

func Solution(A int, B int, K int) int {
	result := 0
	var start int

	if K > A && A != 0 {
		start = K
	} else {
		start = A + A%K
	}

	if start <= B {
		result = (B-start) / K + 1
	}

	return result
}

func TestSolution(t *testing.T)  {
	testCases := []struct{
		A int
		B int
		K int
		ExpectedResult int
	}{
		{6,11,2, 3},
		{0,0,11,1},
	}

	for _, tc := range testCases {
		result := Solution(tc.A, tc.B, tc.K)
		if reflect.DeepEqual(result, tc.ExpectedResult) {
			t.Log("PASS")
		} else {
			t.Errorf("FAIL ")
		}
	}
}
