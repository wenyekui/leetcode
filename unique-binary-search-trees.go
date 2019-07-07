package main

import (
	"fmt"
)

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := n; i >= 2; i-- {
		dp[i] = dp[i-1] + dp[n-i]
	}
}

func main() {
	fmt.Println("")
}
