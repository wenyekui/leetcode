package main

import (
	"fmt"
)

func maxUncrossedLines(A []int, B []int) int {
	M := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		M[i] = make([]int, len(A))
		for j := 0; j < len(A); j++ {
			M[i][j] = -1
		}
	}
	return solve(M, A, B, 0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(M [][]int, A, B []int, i, j int) int {
	if i >= len(A) || j >= len(B) {
		return 0
	}
	if M[i][j] != -1 {
		return M[i][j]
	}
	if A[i] == B[j] {
		M[i][j] = 1 + solve(M, A, B, i+1, j+1)
	} else {
		M[i][j] = max(solve(M, A, B, i+1, j), solve(M, A, B, i, j+1))
	}
	return M[i][j]
}

func main() {

	fmt.Println("")
}
