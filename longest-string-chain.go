package main

import (
	"fmt"
)

func dfs(mem map[string]int, patters map[string][]string, word string) int {
	if count, ok := mem[word]; ok {
		return count
	}

	res := 1
	for i := 0; i <= len(word); i++ {
		p := word[:i] + "*" + word[i:]
		if patters[p] == nil {
			continue
		}

		for _, next := range patters[p] {
			println(word, p, next)
			res = max(res, dfs(mem, patters, next)+1)
		}
	}
	mem[word] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestStrChain(words []string) int {
	patterns := make(map[string][]string)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			p := word[:i] + "*" + word[i+1:]
			if patterns[p] == nil {
				patterns[p] = []string{}
			}
			patterns[p] = append(patterns[p], word)
		}
	}

	// fmt.Println(patterns)

	ans := 0
	mem := make(map[string]int)
	for _, word := range words {
		ans = max(ans, dfs(mem, patterns, word))
	}

	fmt.Println(mem)
	return ans
}

func main() {
	words := []string{
		"ksqvsyq", "ks", "kss", "czvh", "zczpzvdhx", "zczpzvh", "zczpzvhx", "zcpzvh", "zczvh", "gr", "grukmj", "ksqvsq", "gruj", "kssq", "ksqsq", "grukkmj", "grukj", "zczpzfvdhx", "gru",
	}
	ans := longestStrChain(words)
	fmt.Println(ans)
}
