package main

import (
	"check-anagram/v2/ana"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

type anagramFunc func(a, b string) bool

func truncate(s string, n int) string {
	l := len(s)
	if len(s) <= n {
		return s + "(" + strconv.Itoa(l) + ")"
	}
	return s[:n] + "(" + strconv.Itoa(l) + ")"
}

func main() {
	testData := []ana.TestCase{
		{"abc", "abx", false},
		{"abc", "cab", true},
		{"abc", "caba", false},
		{"abcdefghij", "jihgfedcba", true},
		{"the quick brown fox", "xof nworb kciuq eht", true},

		{strings.Repeat("a", 100), strings.Repeat("a", 100), true},
		{strings.Repeat("a", 1000), strings.Repeat("a", 1000), true},
		{strings.Repeat("a", 10000), strings.Repeat("a", 10000), true},
	}

	funcs := map[string]anagramFunc{
		"map":  ana.CheckAnagram,
		"sort": ana.CheckAnagramSort,
	}

	for _, testDatum := range testData {
		for name, fn := range funcs {
			var result bool = false
			r := testing.Benchmark(func(b *testing.B) {
				for _ = range b.N {
					result = fn(testDatum.A, testDatum.B)
				}
			})
			fmt.Printf(" %v %v %v %v Took: %v\n",
				truncate(testDatum.A, 20),
				truncate(testDatum.B, 20),
				result,
				name,
				r.T/time.Duration(r.N),
			)

			if testDatum.Result != result {
				panic(fmt.Sprintf("Failed test: %v", testDatum))
			}

		}

	}

}
