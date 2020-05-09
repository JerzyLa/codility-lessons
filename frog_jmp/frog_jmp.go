package frog_jmp

func Solution(X int, Y int, D int) int {
	if (Y - X) % D != 0 {
		return (Y - X) / D +1
	}

	return (Y - X) / D
}
