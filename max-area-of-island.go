package main

import (
	"fmt"
)

func maxAreaOfIsland(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m, m)
	}
	max := 0

	var dfs func(grid [][]int, i int, j int) int
	dfs = func(grid [][]int, i int, j int) int {
		if i >= n || j >= m || i < 0 || j < 0 {
			return 0
		}
		if visited[i][j] {
			return 0
		}
		visited[i][j] = true
		if grid[i][j] == 0 {
			return 0
		}

		x := 1 + dfs(grid, i+1, j) + dfs(grid, i, j+1) + dfs(grid, i-1, j) + dfs(grid, i, j-1)
		if x > max {
			max = x
		}
		return x
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(grid, i, j)
		}
	}

	return max
}

func main() {
	a := [][]int{
		[]int{1, 1},
		[]int{1, 0},
	}
	fmt.Println(maxAreaOfIsland(a))
}
