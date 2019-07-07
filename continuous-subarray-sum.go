package main

import (
	"fmt"
)

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}

	dp := make([]int, len(nums))
	copy(dp, nums)

	for d := 1; d < len(nums); d++ {
		for i := 0; i < len(nums)-d; i++ {
			dp[i] += nums[i+d]
			if k == 0 {
				if dp[i] == 0 {
					return true
				}
			} else {
				if dp[i]%k == 0 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	nums := []int{23, 2, 4, 6, 7}
	k := 6

	ans := checkSubarraySum(nums, k)
	fmt.Println(ans)

	nums = []int{23, 2, 6, 4, 7}
	k = 6

	ans = checkSubarraySum(nums, k)
	fmt.Println(ans)
}
