package genomic_range_query

import (
	"reflect"
	"testing"
)

var m = map[uint8]int{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

func Solution(S string, P []int, Q []int) []int {
	var result []int
	occurrences := make([][4]int, len(S))

	prev := [4]int{0, 0, 0, 0}
	for i := range occurrences {
		item := m[S[i]]
		prev[item]++
		for j := range prev {
			occurrences[i][j] = prev[j]
		}
	}

	var a [4]int
	zeros := [4]int{0,0,0,0}
	for i := range P {

		if P[i] > 0 {
			a = occurrences[P[i]-1]
		} else {
			a = zeros
		}
		b := occurrences[Q[i]]

		for j := range a {
			if b[j] - a[j] > 0 {
				result = append(result, j+1)
				break
			}
		}
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		S              string
		P              []int
		Q              []int
		ExpectedResult []int
	}{
		{"CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}, []int{2, 4, 1}},
		{"A", []int{0}, []int{0}, []int{1}},
	}

	for _, tc := range testCases {
		result := Solution(tc.S, tc.P, tc.Q)
		if reflect.DeepEqual(result, tc.ExpectedResult) {
			t.Log("PASS")
		} else {
			t.Errorf("FAIL %v, expectes %v", result, tc.ExpectedResult)
		}
	}
}
