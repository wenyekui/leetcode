package main

import (
	"fmt"
)

func possibleBipartition(N int, dislikes [][]int) bool {
}

func main() {
	N := 5
	dislikes := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
		[]int{4, 5},
		[]int{1, 5},
	}
	x := possibleBipartition(N, dislikes)
	fmt.Println(x)
}
