package main

import (
	"fmt"
)

func findNumberOfLIS(nums []int) int {
	N := len(nums)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = []int{1, 1}
	}
	for i := N - 2; i >= 0; i-- {
		for j := i + 1; j < N; j++ {
			if nums[i] < nums[j] {
				if dp[i][0] < dp[j][0]+1 {
					dp[i][0] = dp[j][0] + 1
					dp[i][1] = dp[j][1]
				} else if dp[i][0] == dp[j][0]+1 {
					dp[i][1] += dp[j][1]
				}
			}
		}
	}

	max := 0
	num := 0
	for i := 0; i < N; i++ {
		if dp[i][0] > max {
			max = dp[i][0]
			num = dp[i][1]
		} else if dp[i][0] == max {
			num += dp[i][1]
		}
	}
	return num
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	ans := findNumberOfLIS(nums)
	fmt.Println(ans)

	nums = []int{2, 2, 2, 2, 2}
	ans = findNumberOfLIS(nums)
	fmt.Println(ans)
}
