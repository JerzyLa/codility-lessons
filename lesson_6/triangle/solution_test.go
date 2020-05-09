package triangle

import (
	"reflect"
	"sort"
	"testing"
)

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}

	sort.Ints(A)
	for i:=0; i<len(A)-2; i++ {
		if A[i] + A[i+1] > A[i+2] {
			return 1
		}
	}

	return 0
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A              []int
		ExpectedResult int
	}{
		{[]int{10,2,5,1,8,20}, 1},
		{[]int{10,50,5,1}, 0},
		{[]int{}, 0},
		{[]int{2,0,4}, 0},
		{[]int{1,2}, 0},
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
