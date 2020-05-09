package distinct

import (
	"reflect"
	"testing"
)

func Solution(A []int) int {
	m := map[int]struct{}{}

	result := 0
	for _, el := range A {
		if _, ok := m[el]; !ok {
			m[el] = struct{}{}
			result++
		}
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A              []int
		ExpectedResult int
	}{
		{[]int{2, 1, 1, 2, 3, 1}, 3},
		{[]int{}, 0},
		{[]int{10}, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1},
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
