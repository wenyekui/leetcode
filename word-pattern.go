package main

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	if len(words) != len(pattern) {
		return false
	}
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	for i := 0; i < len(words); i++ {
		c := pattern[i]
		word := words[i]
		_, ok1 := m1[c]
		_, ok2 := m2[word]
		if ok1 || ok2 {
			if m1[c] != word || m2[word] != c {
				return false
			}
		} else {
			m1[c] = word
			m2[word] = c
		}

	}
	return true
}

func main() {
	println(wordPattern("abba", "dog cat cat dog"))
	println(wordPattern("abba", "dog cat cat fish"))
	println(wordPattern("ab", "dog dog"))
}
