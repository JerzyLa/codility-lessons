package dominator

import (
	"reflect"
	"testing"
)

func Solution(A []int) int {
	dominator := len(A)/2 +1
	m := make(map[int]int)

	for i, a := range A {
		m[a]++
		if m[a] == dominator {
			return i
		}
	}

	return -1
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A        []int
		Expected int
	}{
		{[]int{3, 4, 3, 2, 3, -1, 3, 3}, 7},
		{[]int{}, -1},
	}

	for _, tc := range testCases {
		result := Solution(tc.A)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Log("PASS")
		} else {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
