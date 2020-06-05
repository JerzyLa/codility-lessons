package frog_river_one

import (
	"reflect"
	"testing"
)

func Solution(X int, A []int) int {
	m := make(map[int]bool)

	for i, el := range A {
		if _, ok := m[el]; !ok {
			m[el] = true
			if len(m) == X {
				return i
			}
		}
	}

	return -1
}

type TestCase struct {
	A        []int
	X        int
	Expected int
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{[]int{1, 3, 1, 4, 2, 3, 5, 4}, 5, 6},
		{[]int{1, 3}, 3, -1},
		{[]int{1, 2, 3}, 4, -1},
		{[]int{}, 1, -1},
	}

	for _, tc := range testCases {
		result := Solution(tc.X, tc.A)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Logf("PASS")
		} else {
			t.Errorf("FAIL, expected %v but got %v", tc.Expected, result)
		}
	}
}
