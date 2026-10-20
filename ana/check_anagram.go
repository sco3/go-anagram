package ana

func CheckAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := map[int32]int{}
	for _, c := range a {
		m[c]++
	}

	for _, c := range b {
		//fmt.Println(m)
		val := m[c]
		if val == 0 {
			return false
		}
		m[c]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
