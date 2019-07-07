package main

import (
	"fmt"
)

func numSubarrayBoundedMax(A []int, L int, R int) int {
	stack := []int{}
	ans := 0
	for i := 0; i < len(A); i++ {
		left := -1
		for len(stack) != 0 && A[stack[len(stack)-1]] < A[i] {
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) != 0 {
			left = stack[len(stack)-1]
		}

		stack = append(stack, i)

		if A[i] < L || A[i] > R {
			continue
		}
		j := i + 1
		k := i
		for j < len(A) {
			if A[i] < A[j] {
				break
			}
			if A[i] == A[j] {
				ans += (i - left) * (j - k)
				k = j
			}
			j++
		}
		ans += (i - left) * (j - k)
	}
	return ans
}

func main() {
	A := []int{2, 1, 4, 3}
	L := 2
	R := 3

	fmt.Println(numSubarrayBoundedMax(A, L, R))
	A = []int{7, 3, 6, 7, 1}
	L = 1
	R = 5

	//	A := []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}
	//	L := 32
	//	R := 69

	//	A := []int{1, 1, 1}
	//	L := 1
	//	R := 1
	fmt.Println(numSubarrayBoundedMax(A, L, R))
}
