package ana

import "slices"

func CheckAnagramSort(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	ra := []rune(a)
	rb := []rune(b)

	slices.Sort(ra)
	slices.Sort(rb)

	return string(ra) == string(rb)

}
