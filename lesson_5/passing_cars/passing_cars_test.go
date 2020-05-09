package passing_cars

import (
	"reflect"
	"testing"
)

func Solution(A []int) int {
	east := 0
	result := 0
	for _, a := range A {
		if a == 0 {
			east++
		} else {
			result += east
		}
	}

	if result > 1000000000 {
		return -1
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A              []int
		ExpectedResult int
	}{
		{[]int{0,1,0,1,1}, 5},
		{[]int{1}, 0},
		{[]int{0}, 0},
		{[]int{1,0,1,0,1,0}, 3},
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
