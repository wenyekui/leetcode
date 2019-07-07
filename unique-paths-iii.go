package main

import (
	"fmt"
)

type Point [2]int

func uniquePathsIII(grid [][]int) int {
	var start Point

	ans := 0
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				start[0] = i
				start[1] = j
			} else if grid[i][j] == 0 {
				count++
			}
		}
	}

	dfs(grid, make(map[Point]bool), start, &ans, count)
	return ans
}

func dfs(grid [][]int, visited map[Point]bool, current Point, ans *int, count int) {
	x, y := current[0], current[1]
	for _, delta := range [][]int{
		[]int{1, 0},
		[]int{0, 1},
		[]int{-1, 0},
		[]int{0, -1},
	} {
		i := x + delta[0]
		j := y + delta[1]
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
			if grid[i][j] == 2 {
				if len(visited) != count {
					continue
				}
				*ans++
				return
			}

			if grid[i][j] != 0 {
				continue
			}
			p := Point{i, j}
			if visited[p] {
				continue
			}

			visited[p] = true
			dfs(grid, visited, p, ans, count)
			delete(visited, p)
		}
	}
}

func main() {
	grid := [][]int{
		[]int{1, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 2, -1},
	}
	ans := uniquePathsIII(grid)
	fmt.Println(ans)
}
