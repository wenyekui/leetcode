package main

import (
	"fmt"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {
	nRows := len(grid)
	nColumns := len(grid[0])

	rowMax := make([]int, nRows, nRows)
	colomnMax := make([]int, nColumns, nColumns)

	for i := 0; i < nRows; i++ {
		max := 0
		for j := 0; j < nColumns; j++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		rowMax[i] = max
	}

	for j := 0; j < nColumns; j++ {
		max := 0
		for i := 0; i < nRows; i++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		colomnMax[j] = max
	}

	count := 0
	for i := 0; i < nRows; i++ {
		for j := 0; j < nColumns; j++ {
			count += min(rowMax[i], colomnMax[j]) - grid[i][j]
		}
	}
	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	grid := [][]int{
		[]int{3, 0, 8, 4},
		[]int{2, 4, 5, 7},
		[]int{9, 2, 6, 3},
		[]int{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}
