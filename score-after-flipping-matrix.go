package main

import (
	"fmt"
)

func matrixScore(A [][]int) int {
	n := len(A)
	m := len(A[0])

	for i := 0; i < n; i++ {
		if A[i][0] == 0 {
			flipRow(A, m, i)
		}
	}

	for j := 1; j < m; j++ {
		zeros := 0
		ones := 0
		for i := 0; i < n; i++ {
			if A[i][j] == 0 {
				zeros++
			} else {
				ones++
			}
		}
		if zeros > ones {
			flipColumn(A, n, j)
		}
	}
	return score(A, n, m)
}

func score(A [][]int, n, m int) int {
	s := 0
	for i := 0; i < n; i++ {
		t := 0
		for j := 0; j < m; j++ {
			t = t*2 + A[i][j]
		}
		s += t
	}
	return s
}

func flipRow(A [][]int, m, i int) {
	for j := 0; j < m; j++ {
		if A[i][j] == 1 {
			A[i][j] = 0
		} else {
			A[i][j] = 1
		}
	}
}

func flipColumn(A [][]int, n, j int) {
	for i := 0; i < n; i++ {
		if A[i][j] == 1 {
			A[i][j] = 0
		} else {
			A[i][j] = 1
		}
	}
}

func main() {
	A := [][]int{
		[]int{0, 0, 1, 1},
		[]int{1, 0, 1, 0},
		[]int{1, 1, 0, 0},
	}
	fmt.Println(matrixScore(A))
}
