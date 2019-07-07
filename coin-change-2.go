package main

import (
	"fmt"
)

func change(amount int, coins []int) int {
	return helper(amount, 0, map[string]int{}, coins)
}

func helper(amount int, i int, cache map[string]int, coins []int) int {
	key := fmt.Sprintf("%s:%s", amount, i)
	if ans, ok := cache[key]; ok {
		return ans
	}
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}
	if len(coins) == 1 {
		if amount%coins[0] == 0 {
			cache[key] = 1
			return 1
		} else {
			cache[key] = 0
			return 0
		}
	}
	ans := 0
	for j := i; j < len(coins); j++ {
		coin := coins[j]
		if amount >= coin {
			ans += helper(amount-coin, j, cache, coins)
		}
	}
	cache[key] = ans
	return ans
}

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	amount = 3
	coins = []int{2}

	amount = 10
	coins = []int{10}
	amount = 500
	coins = []int{3, 5, 7, 8, 9, 10, 11}
	ans := change(amount, coins)
	fmt.Println(ans)
}
