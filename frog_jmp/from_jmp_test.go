package frog_jmp

import (
	"reflect"
	"testing"
)

type TestCase struct {
	X int
	Y int
	D int
	Expected int
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{10, 85, 30, 3},
	}

	for _, tc := range testCases{
		result := Solution(tc.X, tc.Y, tc.D)

		if reflect.DeepEqual(result, tc.Expected) {
			t.Logf("PASS")
		} else {
			t.Errorf("FAIL, expected %v but got %v", tc.Expected, result)
		}
	}
}
