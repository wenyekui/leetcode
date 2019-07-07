package main

import (
	"fmt"
)

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a

}

func minMoves2(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	avg := sum / len(nums)
	avg2 := avg + 1
	ans := 0
	ans2 := 0
	for _, num := range nums {
		ans += abs(num, avg)
		ans2 += abs(num, avg2)
	}
	return min(ans, ans2)
}

func main() {
	nums := []int{1, 2, 3}
	ans := minMoves2(nums)
	fmt.Println(ans)
}
