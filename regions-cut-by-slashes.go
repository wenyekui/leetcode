package main

import (
	"fmt"
)

func regionsBySlashes(grid []string) int {
	N := len(grid)
	visited := make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, N)
	}
	w
	ans := 0
	for d := 0; d < N-1; d++ {
		for _, pos := range [][]int{
			[]int{0, d}, []int{d, N - 1}, []int{N - 1, 1 + d}, []int{1 + d, 0},
		} {
			i, j := pos[0], pos[1]
			if visited[i][j] || grid[i][j] == ' ' {
				continue
			}
			dfs(grid, visited, i, j, &ans)
		}
	}

	return ans
}

func dfs(grid []string, visited [][]bool, i int, j int, ans *int) {
	visited[i][j] = true
}

func main() {
	fmt.Println("")
}
