package main

import (
	"fmt"
)

func increasingTriplet(nums []int) bool {
	left := 0

	for left < len(nums)-2 && nums[left] >= nums[left+1] {
		left++
	}

	min := left
	mid := left + 1

	for right := mid + 1; mid < len(nums)-1 && right < len(nums); right++ {
		if nums[right] > nums[mid] {
			return true
		}

		if nums[right] < nums[min] {
			min = right
			continue
		}

		if nums[right] > nums[min] && nums[right] < nums[mid] {
			left = min
			mid = right
			continue
		}

		if nums[right] > nums[left] && nums[right] < nums[mid] {
			mid = right
		}

	}
	return false
}

func main() {

	fmt.Println("")
}
