package main

import (
	"fmt"
)

func find(dsu []int, x int) int {
	if dsu[x] != x {
		y := find(dsu, dsu[x])
		dsu[x] = y
	}
	return dsu[x]
}

func union(dsu []int, x int, y int) {
	dsu[find(dsu, x)] = find(dsu, y)
}

func findCircleNum(M [][]int) int {
	N := len(M)
	dsu := make([]int, N)
	for i := 0; i < N; i++ {
		dsu[i] = i
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if M[i][j] == 1 {
				union(dsu, i, j)
			}
		}
	}
	ans := 0
	for x, y := range dsu {
		if x == y {
			ans += 1
		}
	}
	return ans
}

func main() {
	fmt.Println("")
}
