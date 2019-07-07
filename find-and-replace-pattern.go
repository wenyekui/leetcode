package main

import (
	"fmt"
)

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for _, word := range words {
		if len(word) != len(pattern) {
			continue
		}
		i := 0
		m1 := make(map[byte]byte)
		m2 := make(map[byte]byte)
		for i = 0; i < len(word); i++ {
			c1, ok1 := m1[word[i]]
			c2, ok2 := m2[pattern[i]]

			if (!ok1 && ok2) || (ok1 && !ok2) {
				break
			}

			if ok1 && ok2 {
				if word[i] != c2 || pattern[i] != c1 {
					break
				}
			} else {
				m1[word[i]] = pattern[i]
				m2[pattern[i]] = word[i]
			}
		}
		if i == len(word) {
			res = append(res, word)
		}

	}
	return res
}

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words, pattern))
}
