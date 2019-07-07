package main

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W != 0 || len(hand) == 0 {
		return false
	}

	counter := map[int]int{}

	cards := []int{}
	for _, card := range hand {
		if _, ok := counter[card]; !ok {
			cards = append(cards, card)
		}
		counter[card]++
	}

	sort.Ints(cards)
	i := 0
	for i+W-1 < len(cards) {
		k := i
		base := counter[cards[i]]
		for j := i; j < i+W; j++ {
			if counter[cards[j]] < base {
				return false
			}
			if cards[j]-cards[i] != j-i {
				return false
			}
			counter[cards[j]] -= base
			if counter[cards[j]] != 0 && k == i {
				k = j
			}
		}
		if k != i {
			i = k
		} else {
			i = i + W
		}
	}
	if i == len(cards) {
		return true
	}
	return false
}

func main() {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	W := 3

	fmt.Println(isNStraightHand(hand, W))
}
