package odd_occurences

import (
	"reflect"
	"testing"
)

type TestCase struct {
	A []int
	Expected int
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
	}

	for _, tc := range testCases{
		result := Solution(tc.A)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Logf("PASS")
		} else {
			t.Errorf("FAIL, expected %v but got %v", tc.Expected, result)
		}
	}
}
