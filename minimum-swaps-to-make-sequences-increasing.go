package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSwap(A []int, B []int) int {
	N := len(A)
	dp1 := make([]int, N)
	dp2 := make([]int, N)

	dp1[N-1] = 1
	dp2[N-1] = 0

	for i := N - 2; i >= 0; i-- {
		if A[i] < A[i+1] && A[i] < B[i+1] && B[i] < A[i+1] && B[i] < B[i+1] {
			dp2[i] = min(dp1[i+1], dp2[i+1])
			dp1[i] = 1 + dp2[i]
		} else if A[i] < A[i+1] && B[i] < B[i+1] {
			dp2[i] = dp2[i+1]
			dp1[i] = 1 + dp1[i+1]
		} else if A[i] < B[i+1] && B[i] < A[i+1] {
			dp2[i] = dp1[i+1]
			dp1[i] = 1 + dp2[i+1]
		} else {
			dp1[i] = N
			dp2[i] = N
		}
	}
	return min(dp1[0], dp2[0])
}

func main() {
	fmt.Println("")
}
