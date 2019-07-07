package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxWidthRamp(A []int) int {
	N := len(A)
	pos := make([]int, N)
	for i := 0; i < N; i++ {
		pos[i] = i
	}

	sort.Slice(pos, func(i, j int) bool {
		if A[pos[i]] == A[pos[j]] {
			return pos[i] < pos[j]
		}
		return A[pos[i]] < A[pos[j]]

	})

	ans := 0
	m := N
	for _, j := range pos {
		if j > m {
			ans = max(ans, j-m)
		} else {
			m = j
		}
	}
	return ans
}

func main() {

	A := []int{10, 10, 10, 7, 1, 6, 2, 1, 7}
	// A = []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
	ans := maxWidthRamp(A)
	fmt.Println(ans)
}
