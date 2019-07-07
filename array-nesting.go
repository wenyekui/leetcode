package main

import (
	"fmt"
)

func arrayNesting(nums []int) int {
	ans := 0
	used := make([]bool, len(nums))
	for _, num := range nums {
		if used[num] {
			continue
		}
		k := num
		t := 1
		for nums[k] != num {
			used[k] = true
			k = nums[k]
			t++
		}
		if t > ans {
			ans = t
		}
	}
	return ans
}

func main() {
	nums := []int{
		5, 4, 0, 3, 1, 6, 2,
	}
	ans := arrayNesting(nums)
	fmt.Println(ans)
}
