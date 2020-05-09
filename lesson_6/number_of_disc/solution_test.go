package number_of_disc

import (
	"reflect"
	"sort"
	"testing"
)

type Pair struct {
	Key   int
	Value bool
}

type ByKey []Pair

func (s ByKey) Len() int {
	return len(s)
}

func (s ByKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByKey) Less(i, j int) bool {
	if s[i].Key == s[j].Key && s[i].Value != s[j].Value  {
		return s[i].Value == true
	}

	return s[i].Key < s[j].Key
}

func Solution(A []int) int {
	var pairs []Pair

	for i, a := range A {
		pairs = append(pairs, Pair{i - a, true}, Pair{i + a, false})
	}

	sort.Sort(ByKey(pairs))

	activeCircles := 0
	result := 0
	for _, p := range pairs {
		if p.Value {
			result += activeCircles
			activeCircles++
		} else {
			activeCircles--
		}
		if result > 10000000 {
			return -1
		}
	}

	return result
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		A              []int
		ExpectedResult int
	}{
		{[]int{1, 5, 2, 1, 4, 0}, 11},
		{[]int{1, 1, 1}, 3},
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
