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
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func shipWithinDays(weights []int, D int) int {
	N := len(weights)
	dp := make([]int, N)
	copy(dp, weights)

	for i := N - 2; i >= 0; i-- {
		dp[i] += dp[i+1]
	}

	for d := 2; d <= D; d++ {
		for i := 0; i < N-1; i++ {
			sum := 0
			for j := i; j < N-1; j++ {
				sum += weights[j]
				dp[i] = min(dp[i], max(sum, dp[j+1]))
			}
		}
		fmt.Printf("%v\n", dp)
	}
	return dp[0]
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5
	ans := shipWithinDays(weights, D)
	fmt.Println(ans)
}
