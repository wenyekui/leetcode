package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	if N == 1 {
		return nums[0]
	}
	if N == 2 {
		return max(nums[0], nums[1])
	}

	pp := nums[1]

	p := max(nums[2], pp)

	ans := p
	for i := 3; i <= N-1; i++ {
		x := max(nums[i]+pp, p)
		ans = max(ans, x)
		pp = p
		p = x

	}

	if N == 3 {
		return max(nums[0], ans)
	}
	if N == 4 {
		return max(nums[0]+nums[2], ans)
	}

	pp = nums[2]
	p = max(nums[3], pp)

	ans = max(ans, nums[0]+p)

	for i := 4; i <= N-2; i++ {
		x := max(nums[i]+pp, p)
		ans = max(ans, nums[0]+x)
		pp = p
		p = x
	}
	return ans
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	ans := rob(nums)
	fmt.Println(ans)
}
