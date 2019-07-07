package main

import (
	"fmt"
)

func PredictTheWinner(nums []int) bool {
	dp := [20][20]int{}
	sum := [20][20]int{}

	for i, num := range nums {
		dp[i][i] = num
		sum[i][i] = num
	}
	for d := 1; d < len(nums); d++ {
		for i := 0; i < len(nums)-d; i++ {
			sum[i][i+d] = nums[i] + sum[i+1][i+d]
			if dp[i+1][i+d] > dp[i][i+d-1] {
				dp[i][i+d] = sum[i][i+d] - dp[i][i+d-1]
			} else {
				dp[i][i+d] = sum[i][i+d] - dp[i+1][i+d]
			}
		}
	}

	return dp[0][len(nums)-1]*2 >= sum[0][len(nums)-1]
}

func main() {

	fmt.Println("")
}
