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
		{A: "abc", B: "abx", Result: false},
		{A: "abc", B: "cab", Result: true},
		{A: "abc", B: "caba", Result: false},
		{A: "abcdefghij", B: "jihgfedcba", Result: true},
		{A: "the quick brown fox", B: "xof nworb kciuq eht", Result: true},

		{A: strings.Repeat("a", 100), B: strings.Repeat("a", 100), Result: true},
		{A: strings.Repeat("a", 1000), B: strings.Repeat("a", 1000), Result: true},
		{A: strings.Repeat("a", 10000), B: strings.Repeat("a", 10000), Result: true},
	}
	funcs := map[string]anagramFunc{
		"map":  ana.CheckAnagram,
		"sort": ana.CheckAnagramSort,
	}

	for _, testDatum := range testData {
		for name, fn := range funcs {
			result := false
			r := testing.Benchmark(func(b *testing.B) {
				for range b.N {
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
