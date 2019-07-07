package main

import (
	"fmt"
)

const MAX = 1001

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 2
	}

	dp := [MAX][MAX]int{}
	for i := 0; i <= n; i++ {
		dp[n][i] = 0
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= n; j++ {

			if i+j <= n {
				if i+i <= n {
					dp[i][j] = min(dp[i+j][j]+1, dp[i*2][i]+2)
				} else {
					dp[i][j] = dp[i+j][j] + 1
				}

			} else {
				dp[i][j] = MAX
			}

		}
	}

	return dp[2][1] + 2
}

func main() {
	println(minSteps(3))
	fmt.Println("")
}
