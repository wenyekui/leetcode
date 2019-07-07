package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	N := len(nums)
	ans := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				if nums[i]+nums[j] > nums[k] {
					ans++
				} else {
					break
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{}
	fmt.Println("")
}
