package main

import (
	"fmt"
)

// func lengthOfLIS(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	s := 1
// 	prev := nums[0]
// 	max := 1
//
// 	for i := 1; i < len(nums); i++ {
// 		if prev < nums[i] {
// 			s++
// 			if s > max {
// 				max = s
// 			}
// 		} else {
// 			s = 1
// 		}
// 		prev = nums[i]
// 	}
// 	return max
// }

func lengthOfLIS(nums []int) int {

}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
