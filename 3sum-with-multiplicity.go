package main

import (
	"fmt"
)

func threeSumMulti(A []int, target int) int {

	ans := 0
	for i := 1; i < len(A); i++ {

		preSum := make(map[int]int)
		for j := 0; j < i; j++ {
			preSum[A[i]+A[j]]++
		}
		for j := i + 1; j < len(A); j++ {
			ans = (ans + preSum[target-A[j]]) % (1000*1000*1000 + 7)
		}
	}
	return ans
}

func main() {
	A := []int{1, 1, 2, 2, 2, 2}
	target := 5
	A = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	target = 8
	ans := threeSumMulti(A, target)
	fmt.Println(ans)
}
