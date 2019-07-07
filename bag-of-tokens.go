package main

import (
	"fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	ans := 0
	start, end := 0, len(tokens)-1
	for start <= end {
		if tokens[start] <= P {
			P -= tokens[start]
			start++
			ans++
			continue
		}

		if ans > 0 && start < end {
			P += tokens[end]
			end--
			ans--
			continue
		}
		break
	}
	return ans
}

func main() {
	fmt.Println("")
}
