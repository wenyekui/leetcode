package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target = sum / k
	sort.Ints(nums)
	return dfs(&nums, target, k)
}

func dfs(nums []int, target int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	if target == 0 {
		return true
	}
	if nums[len(*nums)-1] > target {
		return false
	}

}

func main() {
	fmt.Println("")
}
