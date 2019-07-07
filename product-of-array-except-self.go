package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	N := len(nums)
	ans := make([]int, N)
	ans[0] = 1
	for i := 1; i < N; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	p := nums[N-1]
	for i := N - 2; i >= 0; i-- {
		ans[i] *= p
		p *= nums[i]
	}
	return ans
}

// 1    2      3             4
//     1    1 * nums[i-1]

// 1   1   1 * 2      1 * 2    1 * 2 * 3

func main() {
	nums := []int{1, 2, 3, 4}
	ans := productExceptSelf(nums)
	fmt.Printf("%v\n", ans)
}
