package main

import (
	"fmt"
)

func subarraysDivByK(A []int, K int) int {
	count := make(map[int]int)
	count[0] = 1
	ans := 0
	preSum := 0
	for _, num := range A {
		preSum += num
		x := preSum
		if x < 0 {
			x = -x
		}

		reminder := preSum % K
		if x < K {
			if preSum > 0 {
				ans += count[x-K]
			} else {
				ans += count[K-x]
			}
		} else {
			ans += count[reminder]
		}
		count[reminder]++
	}
	return ans
}

func main() {
	A := []int{4, 5, 0, -2, -3, 1}
	K := 5

	// A = []int{-1, 2, 9}
	// K = 2

	ans := subarraysDivByK(A, K)
	fmt.Println(ans)
}
