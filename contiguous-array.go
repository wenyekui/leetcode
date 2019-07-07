package main

import (
	"fmt"
)

func findMaxLength(nums []int) int {
	numOne, numZero := 0, 0
	cache := make(map[string]bool)
	for _, num := range nums {
		if num == 1 {
			numOne++
		} else {
			numZero++
		}
	}

	ans := 0
	return ans
}

func main() {
	nums := []int{1, 0}
	ans := findMaxLength(nums)
	fmt.Println(ans)

	nums = []int{1, 0, 0, 1}
	ans = findMaxLength(nums)
	fmt.Println(ans)

	nums = []int{0, 0, 1, 0, 0, 0, 1, 1}
	ans = findMaxLength(nums)
	fmt.Println(ans)

	nums = []int{0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1}
	println(len(nums))
	ans = findMaxLength(nums)
	fmt.Println(ans)

	nums = []int{0, 1, 1, 0, 1, 1, 1, 0}
	ans = findMaxLength(nums)
	fmt.Println(ans)
}
