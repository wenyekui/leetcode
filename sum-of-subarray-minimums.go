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

func sumSubarrayMins(A []int) int {
	N := len(A)
	dp := make([]int, N)
	ans := 0
	for i := 0; i < N; i++ {
		dp[i] = A[i]
		ans = (ans + A[i]) % (1000*1000*1000 + 7)
	}
	for d := 1; d < N; d++ {

		for i := 0; i < N-d; i++ {
			dp[i] = min(dp[i], dp[i+1])
			ans = (ans + dp[i]) % (1000*1000*1000 + 7)
		}
	}

	return ans
}

func main() {
	A := []int{3, 1, 2, 4}
	A = []int{97, 61, 59, 45}
	ans := sumSubarrayMins(A)
	println("------------------")
	fmt.Println(ans)
}
