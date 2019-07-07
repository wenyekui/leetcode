package main

import (
	"fmt"
)

func minimumDeleteSum(s1 string, s2 string) int {
	S := [1001][1001]int{}
	n1 := len(s1)
	n2 := len(s2)
	for i := n1 - 1; i >= 0; i-- {
		S[i][n2] = int(s1[i]) + S[i+1][n2]
	}
	for i := n2 - 1; i >= 0; i-- {
		S[n1][i] = int(s2[i]) + S[n1][i+1]
	}

	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				S[i][j] = S[i+1][j+1]
			} else {
				S[i][j] = min(int(s1[i])+S[i+1][j], int(s2[j])+S[i][j+1])
			}
		}
	}
	return S[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	s1 := "sea"
	s2 := "eat"

	s1 = "delete"
	s2 = "leet"

	s1 = "ccaccjp"
	s2 = "fwosarcwge"

	fmt.Println(minimumDeleteSum(s1, s2))
}
