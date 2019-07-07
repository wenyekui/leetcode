package main

import (
	"fmt"
)

func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}

	dp := [101][101]int{}

	for _, str := range strs {
		nZeros, nOnes := countOnesAndZeros(str)
		for i := m; i >= 0; i-- {
			for j := n; j >= 0; j-- {
				if i >= nZeros && j >= nOnes {
					dp[i][j] = max(dp[i][j], 1+dp[i-nZeros][j-nOnes])
				}
			}
		}
	}
	return dp[m][n]
}

func countOnesAndZeros(str string) (int, int) {
	nOnes := 0
	nZeros := 0
	for _, c := range str {
		if c == '1' {
			nOnes++
		} else {
			nZeros++
		}
	}
	return nZeros, nOnes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func findMaxForm(strs []string, m int, n int) []string {
// 	// fmt.Printf("%v\n", strs)
//
// 	if len(strs) == 0 || (m == 0 && n == 0) {
// 		return []string{}
// 	}
// 	str := strs[0]
// 	nOnes, nZeros := countOnesAndZeros(str)
// 	if nOnes > n || nZeros > m {
// 		return findMaxForm(strs[1:len(strs)], m, n)
// 	}
//
// 	return max(
// 		append(findMaxForm(strs[1:len(strs)], m-nZeros, n-nOnes), str),
//
// 		findMaxForm(strs[1:len(strs)], m, n))
// }

func main() {
	strs := []string{
		"10", "0001", "111001", "1", "0",
	}
	m := 3
	n := 4

	strs = []string{
		"111", "1000", "1000", "1000",
	}

	m = 9
	n = 3

	d := findMaxForm(strs, m, n)
	fmt.Printf("%v\n", d)
	// fmt.Println(d)
}
