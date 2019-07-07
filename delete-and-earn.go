package main

import (
	"fmt"
	"sort"
)

// func deleteAndEarn(nums []int) int {
// 	max := 0
// 	used := map[int]bool{}
// 	for i := 0; i < len(nums); i++ {
// 		num := nums[i]
// 		if used[num] {
// 			continue
// 		}
// 		used[num] = true
//
// 		a1 := []int{}
// 		for j := 0; j < len(nums); j++ {
// 			if nums[j] != num+1 && nums[j] != num-1 && j != i {
// 				a1 = append(a1, nums[j])
// 			}
//
// 		}
// 		x1 := deleteAndEarn(a1)
//
// 		if max < num+x1 {
// 			max = num + x1
// 		}
// 	}
// 	return max
// }
//

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}
	n := len(counter)
	a := make([]int, 0, n+1)
	for num, _ := range counter {
		a = append(a, num)
	}
	sort.Ints(a)
	dp := make([]int, n+1)
	dp[n-1] = a[n-1] * counter[a[n-1]]
	for i := n - 2; i >= 0; i-- {

		if a[i]+1 == a[i+1] {
			dp[i] = max(a[i]*counter[a[i]]+dp[i+2], dp[i+1])
		} else {
			dp[i] = a[i]*counter[a[i]] + dp[i+1]
		}
	}

	return dp[0]
}

func main() {
	nums := []int{3, 4, 2}
	nums = []int{2, 2, 3, 3, 3, 4}
	fmt.Println(deleteAndEarn(nums))

	// fmt.Println(a)
	// fmt.Println(counter)
	// fmt.Println(dp)
}
