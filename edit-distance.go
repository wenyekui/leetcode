package main

import (
	"fmt"
)

func min(arr ...int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < m {
			m = arr[i]
		}
	}
	return m
}

func minDistance(word1 string, word2 string) int {
	N1 := len(word1)
	N2 := len(word2)
	if N1 == 0 {
		return N2
	}
	if N2 == 0 {
		return N1
	}

	dp := make([][]int, N1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, N2)
	}

	if word1[N1-1] != word2[N2-1] {
		dp[N1-1][N2-1] = 1
	}

	for i := N1 - 2; i >= 0; i-- {
		if word1[i] == word2[N2-1] {
			dp[i][N2-1] = N1 - 1 - i
		} else {
			dp[i][N2-1] = min(N1-i, dp[i+1][N2-1]+1)
		}
	}

	for i := N2 - 2; i >= 0; i-- {
		if word1[N1-1] == word2[i] {
			dp[N1-1][i] = N2 - 1 - i
		} else {
			dp[N1-1][i] = min(N2-i, dp[N1-1][i+1]+1)
		}
	}

	for i := N1 - 2; i >= 0; i-- {
		for j := N2 - 2; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = min(dp[i+1][j]+1, dp[i][j+1]+1, dp[i+1][j+1])
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j+1], dp[i+1][j+1]) + 1
			}
		}
	}

	return dp[0][0]
}

func main() {

	word1 := "horse"
	word2 := "ros"

	// word1 = "intention"
	// word2 = "execution"

	ans := minDistance(word1, word2)
	fmt.Println(ans)
}
