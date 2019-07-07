package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	m := make(map[byte]int)
	for _, b := range []byte(s1) {
		m[b]++
	}

	m2 := make(map[byte]int)

	left := 0
	for i := 0; i < len(s2); i++ {
		if m[s2[i]] == 0 {
			left = i + 1
			m2 = make(map[byte]int)
			continue
		}

		m2[s2[i]]++

		for m[s2[i]] < m2[s2[i]] {
			m2[s2[left]]--
			left++
		}

		if i-left+1 == len(s1) {
			return true
		}

	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	// 	s2 = "eidboaoo"
	// 	s1 = "hello"
	// 	s2 = "ooolleoooleh"
	ans := checkInclusion(s1, s2)
	fmt.Println(ans)
}
