package main

import (
	"fmt"
	"sort"
)

func makesquare(nums []int) bool {
	sum := 0
	used := make([]bool, len(nums))
	for _, num := range nums {
		sum += num
	}
	if sum%4 != 0 {
		return false
	}

	side := sum / 4

	sort.Ints(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		if used[i] || nums[i] == side {
			continue
		}
		if nums[i] > side {
			return false
		}
		used[i] = true

		k := side - nums[i]
		for true {

		}
	}
}

func main() {
	fmt.Println("")
}
