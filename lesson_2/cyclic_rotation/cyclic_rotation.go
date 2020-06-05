package cyclic_rotation

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	shift := len(A) - (K % len(A))
	A = append(A, A[:shift]...)

	return A[shift:]
}
