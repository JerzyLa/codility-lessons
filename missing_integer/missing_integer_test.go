package missing_integer

import (
	"reflect"
	"testing"
)

func Solution(A []int) int {
	m := make(map[int]bool)
	for _, el := range A {
		if el > 0 {
			m[el] = true
		}
	}
	for i:=1; i<100002; i++ {
		if _, ok := m[i]; !ok {
			return i
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
		{[]int{1, 3, 6, 4, 1, 2}, 5},
		{[]int{1,2,3}, 4},
		{[]int{-1,-3}, 1},
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

