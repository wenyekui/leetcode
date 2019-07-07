package main

import (
	"fmt"
)

func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	queue := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}

	for len(queue) != 0 {
		i := queue[0][0]
		j := queue[0][1]

		for _, d := range [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0 - 1}} {
			if i+d[0] >= 0 && i+d[0] < n && j+d[1] >= 0 && j+d[1] < m && matrix[i+d[0]][j+d[1]] == -1 {
				matrix[i+d[0]][j+d[1]] = matrix[i][j] + 1
				queue = append(queue, []int{i + d[0], j + d[1]})
			}
		}
		queue = queue[1:]
	}
	return matrix
}

func main() {
	fmt.Println("")
}
