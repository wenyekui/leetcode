package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	n := len(height)

	l := 0
	h := n - 1

	ans := 0
	for l <= h {
		ans = max(ans, (h-l)*min(height[l], height[h]))

		if height[l] < height[h] {
			l++
		} else {
			h--
		}
	}
	return ans
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height = []int{}
	fmt.Println(maxArea(height))
}
