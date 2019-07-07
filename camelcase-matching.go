package main

import (
	"fmt"
)

func camelMatch(queries []string, pattern string) []bool {
	ans := make([]bool, len(queries))
	for i, query := range queries {
		j, k := 0, 0
		for k < len(pattern) {
			if j == len(query) {
				break
			}
			if query[j] == pattern[k] {
				j++
				k++
			} else if query[j] >= 'A' && query[j] <= 'Z' {
				break
			} else {
				j++
			}
		}

		for j < len(query) {
			if query[j] >= 'A' && query[j] <= 'Z' {
				break
			}
			j++
		}

		if k == len(pattern) && j == len(query) {
			ans[i] = true
		}
	}
	return ans
}

func main() {
	fmt.Println("")
}
