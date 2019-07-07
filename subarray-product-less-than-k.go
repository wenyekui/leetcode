package main

import (
	"fmt"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	N := len(nums)

	i := 0
	p := 1
	ans := 0
	for j := 0; j < N; j++ {
		p = p * nums[j]
		for p >= k && i < j {
			p = p / nums[i]
			i++
		}
		if p < k {
			ans++
		}
	}
	return N*(N+1)/2 - ans
}

func main() {
	fmt.Println("")
}
