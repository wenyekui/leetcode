package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
	N := len(prices)
	sell := make([]int, N)
	buy := make([]int, N)

	sell[N-1] = prices[N-1]
	for i := N - 2; i >= 0; i-- {
		buy[i] = max(buy[i+1], sell[i+1]-prices[i]-fee)
		sell[i] = max(buy[i+1]+prices[i], sell[i+1])
	}
	return buy[0]
}

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	ans := maxProfit(prices, fee)
	fmt.Println(ans)
}
