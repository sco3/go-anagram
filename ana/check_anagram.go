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
		val--
		if val == 0 {
			delete(m, c)
		} else {
			m[c] = val
		}
	}

	if len(m) == 0 {
		return true
	}
	return false
}
