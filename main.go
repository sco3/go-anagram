package main

import "fmt"

func main() {
	checkAnagramAndPrint("abc", "abx")
	checkAnagramAndPrint("abc", "cab")
	checkAnagramAndPrint("abc", "caba")
}

func checkAnagramAndPrint(a string, b string) {
	fmt.Println(a, " ", b, " ", checkAnagram(a, b))

}

func checkAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := map[int32]int{}
	for _, c := range a {
		m[c]++
	}

	for _, c := range b {
		fmt.Println(m)
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
