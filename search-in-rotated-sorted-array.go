package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	i, j := 0, len(nums)-1
	for i < j {
		m := (i + j) / 2
		if nums[i] < target {
			if nums[m] > target {
				j = m
			} else {
				i = m
			}
		} else {
			if nums[m] > nums[i] {
				i = m
			} else {
				if nums[m] > target {
					j = m
				} else {

				}
			}
		}
	}
	if nums[i] == target {
		return i
	}
	return -1
}

func main() {
	fmt.Println("")
}
