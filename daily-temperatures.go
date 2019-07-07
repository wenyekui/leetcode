package main

import (
	"fmt"
)

func dailyTemperatures(T []int) []int {
	n := len(T)
	ans := make([]int, n)
	if n == 0 {
		return ans
	}
	ans[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		if T[i] < T[i+1] {
			ans[i] = 1
			continue
		}

		k := ans[i+1]
		for true {

			if T[i] < T[i+k+1] {
				ans[i] = 1 + k
				break
			}

			if ans[i+k+1] == 0 {
				ans[i] = 0
				break
			}

			k += ans[i+k+1]
		}
	}
	return ans
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	T = []int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}
	T = []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(T)
	fmt.Println(ans)
}
