package main

import (
	"fmt"
)

func min(nums ...int) int {
	t := nums[0]
	for _, x := range nums {
		if t > x {
			t = x
		}
	}
	return t
}

func minFallingPathSum(A [][]int) int {
	n := len(A)

	dp := make([][]int, n, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+2, n+2)
		dp[i][0] = 10000
		dp[i][n+1] = 10000
	}

	for j := 1; j <= n; j++ {
		dp[n-1][j] = A[n-1][j-1]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 1; j <= n; j++ {
			dp[i][j] = min(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1]) + A[i][j-1]
		}
	}

	return min(dp[0]...)
}

func main() {
	fmt.Println("")
}
