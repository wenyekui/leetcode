package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	ans := 0
	diff := 0
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			diff++
			continue
		}
		if diff != 0 {
			x := A[i] - A[i-1]
			if diff >= x {
				ans += x * (x + 1) / 2
				diff = diff - x + 1
			} else {
				ans += diff * (diff + 1) / 2
				diff = 0
			}
		}
	}
	ans += diff * (diff + 1) / 2

	return ans
}

func main() {
	minIncrementForUnique([]int{1, 2, 3})
	fmt.Println("")
}
