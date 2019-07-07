package main

import (
	"fmt"
)

func numMatchingSubseq(S string, words []string) int {
	m := make(map[byte][]int)

	for i, c := range []byte(S) {
		if m[c] == nil {
			m[c] = []int{}
		}
		m[c] = append(m[c], i)
	}
	ans := 0
	for _, word := range words {
		start := -1
		count := 0
		for _, c := range []byte(word) {
			if m[c] == nil {
				break
			}

			for _, i := range m[c] {
				if i > start {
					start = i
					count++
					break
				}
			}

		}
		if count == len(word) {
			ans++
		}

	}
	return ans
}

func main() {
	S := "abcde"
	words := []string{
		"a", "bb", "acd", "ace",
	}
	ans := numMatchingSubseq(S, words)
	fmt.Println(ans)
}
