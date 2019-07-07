package main

import (
	"fmt"
	"sort"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	sort.Ints(nums)
	return dfs(nums, target)
}

func dfs(nums []int, target int) bool {
	fmt.Printf("%v %d\n", nums, target)
	println("----------------")

	if len(nums) == 0 {
		if target == 0 {
			return true
		}
		return false
	}
	if nums[len(nums)-1] > target {
		return false
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] > target-nums[len(nums)-1] {
		i--
	}

	return dfs(nums[0:i+1], target-nums[len(nums)-1]) || dfs(nums[0:len(nums)-1], target)
}

func main() {
	nums := []int{1, 5, 11, 5}
	var ans bool
	ans = canPartition(nums)
	fmt.Println(ans)

	nums = []int{1, 2, 3, 5}
	ans = canPartition(nums)
	fmt.Println(ans)

	nums = []int{
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}

	//	nums = []int{99, 2, 3, 98}
	//	ans = canPartition(nums)

	nums = []int{2, 2, 3, 5}
	ans = canPartition(nums)
	fmt.Println(ans)
}
