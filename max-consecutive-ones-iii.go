package main

import (
	"fmt"
)

func longestOnes(A []int, K int) int {
	ans := 0
	i, j, l := 0, 0, 0
	for j = 0; j < len(A); j++ {
		if A[j] == 1 {
			l++
		} else {
			if K == 0 {
				for i < j {
					if A[i] == 0 {
						K++
					}
					l--
					i++
					if K > 0 {
						break
					}
				}
			}
			if K > 0 {
				K--
				l++
			}
		}
		if l > ans {
			ans = l
		}
	}

	return ans
}

func main() {
	A := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K := 2

	A = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	K = 3
	ans := longestOnes(A, K)
	fmt.Println(ans)
}
