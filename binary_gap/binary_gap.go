package binary_gap

import (
	"strconv"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Solution(N int) int {
	str := strconv.FormatInt(int64(N), 2)

	result := 0
	gap := -1
	for _, char := range str{
		if char == '1' {
			result = Max(result, gap)
			gap = 0
		} else if gap != -1 {
			gap++
		}
	}

	return result
}
