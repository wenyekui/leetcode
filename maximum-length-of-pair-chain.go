package main

import (
	"fmt"
	"sort"
)

func main() {
	pairs := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
	}
	fmt.Println(findLongestChain(pairs))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	M := make([]int, n, n)

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	M := 1
	m := n - 1

	for i := n - 2; i >= 0; i-- {
		if pairs[i][1] < pairs[m][0] {
			M[i] = M[i+1] + 1
			m = i
		} else {
			M[i] = M[i+1]
		}
	}

	return M[0]
}
