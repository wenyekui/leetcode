package main

import (
	"fmt"
)

func smallestRangeII(A []int, K int) int {
	min := 10001
	max := -1

	sum := 0
	for _, a := range A {
		sum += a
	}

	avg := sum / len(A)

	for _, a := range A {
		var b int
		if a > avg {
			b = a - K
		} else {
			b = a + K
		}
		if b < min {
			min = b
		}
		if b > max {
			max = b
		}
	}
	return max - min
}

func main() {
	x := smallestRangeII([]int{1, 3, 6}, 3)
	x = smallestRangeII([]int{0, 10}, 2)
	x = smallestRangeII([]int{2, 7, 2}, 1)
	fmt.Println(x)
}
