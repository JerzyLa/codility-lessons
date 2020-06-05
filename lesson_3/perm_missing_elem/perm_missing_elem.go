package perm_missing_elem

func Solution(A []int) int {
	perm := make([]bool, len(A) + 1)

	for _, el := range A {
		perm[el-1] = true
	}
	for i, el := range perm {
		if el == false {
			return i+1
		}
	}

	return -1
}
