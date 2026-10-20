package main

import (
	"check-anagram/v2/ana"
	"fmt"
)

type anagramFunc func(a, b string) bool

func main() {
	testData := []ana.TestCase{
		{"abc", "abx", false},
		{"abc", "cab", true},
		{"abc", "caba", false},
	}

	funcs := map[string]anagramFunc{
		"map":  ana.CheckAnagram,
		"sort": ana.CheckAnagramSort,
	}

	for _, testDatum := range testData {
		for name, fn := range funcs {
			var result bool = false
			result = fn(testDatum.A, testDatum.B)
			fmt.Println(testDatum.A, " ", testDatum.B, " ", result, name)
			if testDatum.Result != result {
				panic(fmt.Sprintf("Failed test: %v", testDatum))
			}
		}

	}

}
