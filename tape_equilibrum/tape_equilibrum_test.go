package tape_equilibrum

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
		{[]int{3,1,2,4,3}, 1},
		{[]int{}, 0},
		{[]int{3}, 3},
		{[]int{3, 4, 5}, 2},
		{[]int{-1000, 1000}, 2000},
		{[]int{-2, -3, -4, -1}, 0},
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
