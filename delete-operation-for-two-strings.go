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

func minDistance(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
	dp := make([][]int, n1) [501][501]int{}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return n1 + n2 - dp[n1][n2]*2
}

func main() {
	ans := minDistance("sea", "eat")
	fmt.Println(ans)
}
