package main

import (
	"fmt"
)

// func maxProduct(words []string) int {
// 	n := len(words)
// 	masks := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		for _, c := range words[i] {
// 			masks[i] |= (1 << (byte(c) - 'a'))
// 		}
// 	}
// 	fmt.Println(masks)
// 	max := 0
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			if (masks[i] & masks[j]) == 0 {
// 				t := len(words[i]) * len(words[j])
// 				if t > max {
// 					max = t
// 				}
// 			}
// 		}
// 	}
// 	return max
// }

func maxProduct(words []string) int {
	n := len(words)
	masks := make([]uint, n)
	max := 0
	for i := 0; i < n; i++ {
		for _, c := range []byte(words[i]) {
			masks[i] |= (1 << (c - 'a'))
		}
		for j, word2 := range words[:i] {
			if masks[i]&masks[j] == 0 {
				t := len(words[i]) * len(word2)
				if t > max {
					max = t
				}
			}
		}
	}

	return max
}

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}

	fmt.Println(maxProduct(words))
}
