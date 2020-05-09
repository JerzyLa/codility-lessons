package perm_missing_elem

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
		{[]int{2,3,1,5}, 4},
		{[]int{}, 1},
		{[]int{1}, 2},
		{[]int{2}, 1},
		{[]int{2,3,4,5}, 1},
		{[]int{1,2,3,4,5}, 6},
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
