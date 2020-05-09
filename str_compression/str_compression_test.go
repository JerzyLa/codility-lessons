package str_compression_test

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Solution(A string) string {
	if A == "" {
		return ""
	}

	var result []string
	prev := A[0]
	num := 1
	for _, c := range A[1:] {
		if prev == uint8(c) {
			num++
		} else {
			if num > 1 {
				result = append(result, []string{strconv.Itoa(num), string(prev)}...)
				num = 1
			} else {
				result = append(result, string(prev))
			}
		}
		prev = uint8(c)
	}

	if num > 1 {
		result = append(result, []string{strconv.Itoa(num), string(prev)}...)
		num = 1
	} else {
		result = append(result, string(prev))
	}

	return strings.Join(result, "")
}

func TestSolution(t *testing.T) {
	testCases := []struct{
		A string
		Expected string
	}{
		{"", ""},
		{"AAABBB", "3A3B"},
		{"A", "A"},
		{"ABAABCCCDCCC", "AB2AB3CD3C"},
	}

	for _, tc := range testCases{
		result := Solution(tc.A)
		if reflect.DeepEqual(result, tc.Expected) {
			t.Logf("PASS")
		} else {
			t.Errorf("FAIL, result was %v but expected %v", result, tc.Expected)
		}
	}
}
