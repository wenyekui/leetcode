package main

import (
	"fmt"
)

// func subarraySum(nums []int, k int) int {
// 	r := 0
// 	for i := 0; i < len(nums); i++ {
// 		sum := 0
// 		for j := i; j < len(nums); j++ {
// 			sum += nums[j]
// 			if sum == k {
// 				r++
// 			}
// 		}
// 	}
// 	return r
// }

func subarraySum(nums []int, k int) int {
	preSum := make(map[int]int)
	preSum[0] = 1
	sum := 0
	res := 0
	for _, num := range nums {
		sum += num
		val, ok := preSum[sum-k]
		if ok {
			res += val
		}
		preSum[sum] += 1
	}
	return res
}

func main() {
	nums := []int{1, 1, 1}
	k := 2
	x := subarraySum(nums, k)
	fmt.Println(x)
}
