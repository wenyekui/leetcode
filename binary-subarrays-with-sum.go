package main

import (
	"fmt"
)

func numSubarraysWithSum(A []int, S int) int {

	preSum := make(map[int]int)
	ans := 0
	sum := 0
	preSum[0] = 1
	for _, bit := range A {
		sum += bit
		ans += preSum[sum-S]
		preSum[sum] += 1
	}
	return ans
}

func main() {
	A := []int{1, 0, 1, 0, 1}
	S := 2

	// A = []int{0, 0, 0, 0, 0}
	// S = 0

	ans := numSubarraysWithSum(A, S)
	fmt.Println(ans)
}
