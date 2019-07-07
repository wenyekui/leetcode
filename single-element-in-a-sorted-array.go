package main

import (
	"fmt"
)

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	mid := len(nums) / 2
	if nums[mid-1] == nums[mid] {
		return singleNonDuplicate(nums[0 : mid-1])
	}

	if nums[mid+1] == nums[mid] {
		return singleNonDuplicate(nums[mid+2 : len(nums)])

	}
	return nums[mid]
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}
