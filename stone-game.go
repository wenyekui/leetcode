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

// dp[i] = max(p[i] - dp[i + 1], p[i + d] - dp[i])

func stoneGame(piles []int) bool {
	n := len(piles)

	S := [500][500]int{}
	M := [500][500]int{}

	for i := 0; i < n; i++ {
		S[i][i] = piles[i]
		M[i][i] = piles[i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			M[j][j+i] = M[j][j+i-1] + piles[j+i]
			S[j][j+i] = M[j][j+i] - min(S[j+1][j+i], S[j][j+i-1])
		}
	}
	return S[0][n-1] > M[0][n-1]/2
}

// func stoneGame(piles []int) bool {
// 	sum := 0
// 	for i := 0; i < len(piles); i++ {
// 		sum += piles[i]
// 	}
// 	return highest(sum, piles) > (sum / 2)
// }
//
// func highest(sum int, piles []int) int {
// 	n := len(piles)
// 	if n == 0 {
// 		return 0
// 	}
//
// 	s1 := highest(sum-piles[0], piles[1:n])
// 	s2 := highest(sum-piles[n-1], piles[0:n-1])
// 	if s1 > s2 {
// 		return sum - s2
// 	} else {
// 		return sum - s1
// 	}
// }

func main() {
	fmt.Println("")
}
