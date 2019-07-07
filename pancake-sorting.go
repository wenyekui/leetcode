package main

import (
	"fmt"
)

func pancakeSort(A []int) []int {
	n := len(A)
	ans := []int{}
	j := n - 1
	for n > 0 {
		for i := 0; i < n-1; i++ {
			if A[i] == n {
				j = i
				if i != 0 {
					ans = append(ans, i)
				}
				ans = append(ans, n)
				break
			}
		}
		if j != n-1 {
			for i := 0; i < j/2+1; i++ {
				A[j-i], A[i] = A[i], A[j-i]
			}
			for i := 0; i < n/2; i++ {
				A[n-1-i], A[i] = A[i], A[n-1-i]
			}
		}
		n--
	}
	return ans
}

func main() {
	A := []int{3, 2, 4, 1}
	ans := pancakeSort(A)
	fmt.Println(ans)
}
