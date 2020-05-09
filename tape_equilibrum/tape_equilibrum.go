package tape_equilibrum

func sum(input ...int) int {
	result := 0
	for _, el := range input {
		result += el
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Solution(A []int) int {
	if len(A) == 0{
		return 0
	}
	if len(A) == 1 {
		return A[0]
	}

	all := sum(A...)

	part := A[0]
	result := abs(all - 2*part)
	for _, el := range A[1:len(A)-1] {
		part += el
		diff := abs(all - 2*part)
		if diff < result {
			result = diff
		}
	}

	return result
}
