package equi_leader

import (
	"reflect"
	"testing"
)

func Solution(A []int) int {
	half := len(A) / 2
	m := make(map[int]int)
	var leader int

	for _, a := range A {
		m[a]++
		if m[a] > half {
			leader = a
		}
	}

	res := 0
	allLeadersCount := m[leader]
	leadersCount := 0
	for i, a := range A {
		if a == leader {
			leadersCount++
		}
		if (leadersCount > (i+1)/2) && ((allLeadersCount - leadersCount) > (len(A)-i-1)/2) {
			res++
		}
	}

	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A        []int
		Expected int
	}{
		{[]int{4, 3, 4, 4, 4, 2}, 2},
		{[]int{4, 4, 2, 5, 3, 4, 4, 4}, 3},
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
