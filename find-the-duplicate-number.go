package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < 32; i++ {
		bit := 1 << uint(i)
		c1, c2 := 0, 0
		for j := 0; j < n; j++ {
			if bit&j > 0 {
				c1++
			}
			if bit&nums[j] > 0 {
				c2++
			}
		}
		if c2 > c1 {
			ans += bit
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 1}
	fmt.Println(findDuplicate(nums))
}
