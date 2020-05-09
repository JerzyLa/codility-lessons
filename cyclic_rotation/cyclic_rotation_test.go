package cyclic_rotation

import (
	"reflect"
	"testing"
)

type TestCase struct {
	A []int
	K int
	Expected []int
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
		{[]int{0, 0, 0}, 1, []int{0, 0, 0}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{}, 4, []int{}},
	}

	for _, tc := range testCases{
		result := Solution(tc.A, tc.K)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Logf("PASS")
		} else {
			t.Errorf("FAIL, expected %v but got %v", tc.Expected, result)
		}
	}
}
