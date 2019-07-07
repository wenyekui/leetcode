package main

import (
	"fmt"
)

func main() {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	x := maximalSquare(matrix)
	fmt.Println(x)
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix)
	m := len(matrix[0])
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != '1' {
				continue
			}
			l := 1
			for i+l < n && j+l < m {
				if matrix[i+l][j+l] != '1' {
					break
				}
				flag := true
				for x := 0; x < l; x++ {
					if matrix[i+l][j+x] != '1' || matrix[i+x][j+l] != '1' {
						flag = false
						break
					}
				}
				if !flag {
					break
				}
				l++
			}
			if l > max {
				max = l
			}
		}
	}
	return max * max
}
