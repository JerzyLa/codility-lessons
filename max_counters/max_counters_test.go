package max_counters

import (
	"reflect"
	"testing"
)

func Solution(N int, A []int) []int {
	result := make([]int, N)

	max := 0
	maxCounterVal := 0
	for _, el := range A {
		if el <= len(result) {
			if result[el-1] < maxCounterVal {
				result[el-1] = maxCounterVal
			}

			result[el-1]++
			if result[el-1] > max {
				max = result[el-1]
			}
		} else {
			maxCounterVal = max
		}
	}

	for i := range result {
		if result[i] < maxCounterVal {
			result[i] = maxCounterVal
		}
	}

	return result
}

type TestCase struct {
	A []int
	N int
	Expected []int
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{[]int{3,4,4,6,1,4,4}, 5, []int{3,2,2,4,2}},
	}

	for _, tc := range testCases{
		result := Solution(tc.N, tc.A)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Logf("PASS")
		} else {
			t.Errorf("FAIL, expected %v but got %v", tc.Expected, result)
		}
	}
}
