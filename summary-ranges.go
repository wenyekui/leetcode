package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	ans := []string{}
	if len(nums) == 0 {
		return ans
	}
	start := 0
	end := 0
	for end = 0; end < len(nums); end++ {
		if end == len(nums)-1 || nums[end]+1 != nums[end+1] {
			if start == end {
				ans = append(ans, fmt.Sprintf("%d", nums[start]))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[end]))
			}
			start = end + 1
		}
	}
	return ans
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	ans := summaryRanges(nums)
	fmt.Printf("%s\n", ans)

	nums = []int{0, 2, 3, 4, 6, 8, 9}
	ans = summaryRanges(nums)
	fmt.Printf("%s\n", ans)
}
