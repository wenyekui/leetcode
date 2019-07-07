package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func canReorderDoubled(A []int) bool {
	counter := make(map[int]int)

	for _, a := range A {
		counter[a]++
	}
	keys := make([]int, 0, len(A))
	for k, _ := range counter {
		keys = append(keys, k)
	}

	for _, k := range keys {
		if counter[k] == 0 {
			continue
		}
		if k == 0 {
			if counter[k]%2 != 0 {
				return false
			}
			counter[k] = 0
			continue
		}
		x := k
		for x%2 == 0 {
			x = x / 2
		}
		for abs(x) <= abs(k) {

			if counter[x] != 0 {
				if counter[x] > counter[2*x] {
					return false
				}
				counter[2*x] -= counter[x]
				counter[x] = 0
			}
			x *= 2
		}
	}
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	A := []int{4, -2, 2, -4}
	A = []int{0, 0, 0, 1}
	ans := canReorderDoubled(A)
	fmt.Println(ans)
}
