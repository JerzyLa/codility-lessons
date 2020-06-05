package odd_occurences

func Solution(A []int) int {
	m := make(map[int]bool)

	for _, el := range A {
		if _, ok := m[el]; !ok {
			m[el] = false
		} else {
			m[el] = !m[el]
		}
	}

	for k, v := range m {
		if v == false {
			return k
		}
	}

	return 0
}
