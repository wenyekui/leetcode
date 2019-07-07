package main

import (
	"fmt"
)

func minDeletionSize(A []string) int {
	if len(A) <= 1 {
		return 0
	}
	ans := 0

	for j := 0; j < len(A[0]); j++ {
		flag := true
		for i := 0; i < len(A)-1; i++ {
			if A[i][j] > A[i+1][j] {
				ans++
				flag = false
				break
			} else if A[i][j] == A[i+1][j] {
				flag = false
			}
		}
		if flag {
			return ans
		}
	}
	return ans
}

func main() {
	A := []string{"yc", "yb", "za"}
	ans := minDeletionSize(A)
	fmt.Println(ans)
}
