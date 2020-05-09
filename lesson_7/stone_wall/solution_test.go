package stone_wall

import (
	"reflect"
	"testing"
)

func Solution(H []int) int {
	if len(H) == 0 {
		return 0
	}

	stack := []int{H[0]}
	res := 1
	prev := H[0]
	for _, h := range H[1:] {
		if h > prev {
			stack = append(stack, h)
			res++
		} else if h < prev {
			for i := len(stack) - 1; i>=0 && stack[i] > h; i-- {
				stack = stack[:i]
			}
			if len(stack) == 0 || stack[len(stack)-1] != h {
				stack = append(stack, h)
				res++
			}
		}

		prev = h
	}

	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		H        []int
		Expected int
	}{
		{[]int{8, 8, 5, 7, 9, 8, 7, 4, 8}, 7},
		{[]int{8,8,8,8,8}, 1},
		{[]int{8,8,8,9,8,8,8}, 2},
		{[]int{8,8,8,7,8,8,8}, 3},
	}

	for _, tc := range testCases {
		result := Solution(tc.H)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Log("PASS")
		} else {
			t.Errorf("FAIL: got %v, expectes %v", result, tc.Expected)
		}
	}
}
