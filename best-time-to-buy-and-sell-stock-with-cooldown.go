package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	M := make([]int, len(prices), len(prices))
	for i := 0; i < len(prices); i++ {
		M[i] = -1
	}
	buy(prices, 0, M)
	return M[0]
}

func buy(prices []int, i int, M []int) int {
	if i >= len(prices) {
		return 0
	}
	if M[i] != -1 {
		return M[i]
	}

	M[i] = max(buy(prices, i+1, M), sell(prices[i], prices, i+1, M))
	return M[i]
}

func sell(buy_price int, prices []int, i int, M []int) int {
	if i >= len(prices) {
		return -buy_price
	}
	return max(prices[i]-buy_price+buy(prices, i+2, M), sell(buy_price, prices, i+1, M))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))

}
