package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minFlipsMonoIncr(S string) int {
	N := len(S)
	nZeros := make([]int, N)
	nZeros[N-1] = int('1' - S[N-1])

	for i := N - 2; i >= 0; i-- {
		nZeros[i] = int('1'-S[i]) + nZeros[i+1]
	}

	dp := make([]int, N)
	dp[N-1] = 0
	for i := N - 2; i >= 0; i-- {
		dp[i] = min(nZeros[i], dp[i+1]+int(S[i]-'0'))
	}
	return dp[0]
}

func main() {
	ans := minFlipsMonoIncr("00110")
	fmt.Println(ans)

	ans = minFlipsMonoIncr("00011000")
	fmt.Println(ans)
}
