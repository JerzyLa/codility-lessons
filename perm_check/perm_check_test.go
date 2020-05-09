package perm_check

import (
	"reflect"
	"testing"
)

func Solution(A []int) int {
	perm := make([]bool, len(A))
	for _, el := range A {
		if el > len(A) {
			return 0
		} else if perm[el-1] == true {
			return 0
		} else {
			perm[el-1] = true
		}
	}

	for _, el := range perm{
		if el != true {
			return 0
		}
	}

	return 1
}

type TestCase struct {
	A []int
	Expected int
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{[]int{4,1,3,2}, 1},
		{[]int{4,1,3}, 0},
		{[]int{1}, 1},
		{[]int{2}, 0},
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

