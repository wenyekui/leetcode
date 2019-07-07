package main

import (
	"fmt"
)

func isIdealPermutation(A []int) bool {
	if len(A) == 0 {
		return true
	}
	max := A[0]
	for i := 1; i < len(A)-1; i++ {
		if A[i+1] < max {
			return false
		}
		if A[i] > max {
			max = A[i]
		}
	}
	return true
}

func main() {
	fmt.Println("")
}
