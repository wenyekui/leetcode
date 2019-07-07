package main

import (
	"fmt"
)

func eventualSafeNodes(graph [][]int) []int {
	N := len(graph)
	M := make([]int, N)
	pre := make(map[int]bool)

	for i := 0; i < N; i++ {
		pre[i] = true
		bt(graph, i, pre, M)
		pre[i] = false
	}
	ans := []int{}
	for i := 0; i < N; i++ {
		if M[i] == 1 {
			ans = append(ans, i)
		}
	}
	return ans
}

func bt(graph [][]int, i int, pre map[int]bool, M []int) int {

	for _, x := range graph[i] {
		if M[x] == 1 {
			continue
		}
		if M[x] == -1 {
			return -1
		}
		if pre[x] {
			M[x] = -1
			return -1
		}
		pre[x] = true
		y := bt(graph, x, pre, M)
		pre[x] = false
		if y == -1 {
			M[x] = -1
			return -1
		}
	}
	M[i] = 1
	return 1
}

func main() {
	fmt.Println("")
}
