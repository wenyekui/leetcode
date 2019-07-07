package main

import (
	"fmt"
)

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return []int{-1}
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = -2
	}

	max := nums[0]
	maxIndex := 0
	n := len(nums)

	for i, num := range nums {
		if num > max {
			max = num
			maxIndex = i
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] == max {
			ans[i] = -1
		}
	}

	for i := 1; i < len(nums); i++ {
		j := ((maxIndex-i)%n + n) % n
		if ans[j] == -1 {
			continue
		}
		if ans[(j+1)%n] == -1 {
			ans[j] = maxIndex
			continue
		}
		if nums[(j+1)%n] > nums[j] {
			ans[j] = (j + 1) % n
			continue
		}
		k := (j + 1) % n
		for ans[k] != -1 {
			if nums[j] < nums[k] {
				ans[j] = k
				break
			}
			k = ans[k]

		}
		if ans[k] == -1 {
			ans[j] = maxIndex
		}

	}

	for i, v := range ans {
		if v != -1 {
			ans[i] = nums[v]
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 1}
	nums = []int{100, 1, 11, 1, 120, 111, 123, 1, -1, -100}
	nums = []int{1, 1, 1, 1, 1, 1}
	ans := nextGreaterElements(nums)
	fmt.Printf("%v\n", ans)
}
