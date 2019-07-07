package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func lastStoneWeightII(stones []int) int {
	fmt.Println(stones)

	if len(stones) == 1 {
		return stones[0]
	}

	if len(stones) == 2 {
		return abs(stones[0], stones[1])
	}

	x := stones[0]
	m := 101
	for _, s := range stones[1:] {
		m = min(m, lastStoneWeightII(append(stones[2:], abs(x, s))))
	}
	println(m)
	return m
}

// func lastStoneWeightII(stones []int) int {
// 	N := 100
// 	buckets := make([]int, N+1)
// 	for _, s := range stones {
// 		buckets[s]++
// 	}
//
// 	i := N
// 	j := 1
// 	for i > j {
// 		for i >= 1 && buckets[i] == 0 {
// 			i--
// 		}
//
// 		for j <= N && buckets[j] == 0 {
// 			j++
// 		}
//
// 		if i == j {
// 			break
// 		}
//
// 		println("test", i, j)
//
// 		if buckets[i] > buckets[j] {
// 			buckets[i] = buckets[i] - buckets[j]
// 			if i-j != j {
// 				buckets[i-j] += buckets[j]
// 				buckets[j] = 0
// 			}
// 		} else {
// 			buckets[j] = buckets[j] - buckets[i]
// 			if i-j != i {
// 				buckets[i-j] += buckets[i]
// 				buckets[i] = 0
// 			}
// 		}
// 		if i-j < j {
// 			j = i - j
// 		}
//
// 	}
// 	if i <= 0 {
// 		return 0
// 	}
// 	if buckets[i]%2 == 0 {
// 		return 0
// 	}
// 	return i
// }

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	stones = []int{2, 1, 2, 2, 2}
	stones = []int{4, 4, 1, 1, 1}
	stones = []int{31, 26, 33, 21, 40}
	ans := lastStoneWeightII(stones)
	fmt.Println(ans)
}
