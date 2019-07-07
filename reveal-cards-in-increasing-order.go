package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	ans := []int{}
	n := len(deck)
	if n == 0 {
		return ans
	}
	sort.Ints(deck)
	ans = append(ans, deck[n-1])
	for i := n - 2; i >= 0; i-- {
		ans = append(ans, ans[0])
		ans = ans[1:len(ans)]
		ans = append(ans, deck[i])
	}
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return ans
}

func main() {
	deck := []int{17, 13, 11, 2, 3, 5, 7}
	fmt.Println(deckRevealedIncreasing(deck))
}
