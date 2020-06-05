package binary_gap

import (
	"testing"
)

func TestSolution(t *testing.T) {
	testInputs := []int{9, 529, 1041}
	expectedResults := []int{2, 4, 5}

	for i := 0; i < len(testInputs); i++ {
		result := Solution(testInputs[i])

		if result == expectedResults[i] {
			t.Logf("PASS")
		} else {
			t.Errorf("FAIL, expected %v but got %v", expectedResults[i], result)
		}
	}
}
