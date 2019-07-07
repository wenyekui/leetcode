package main

import (
	"fmt"
)

func longestMountain(A []int) int {
	k := 0
	m := 0
	i := 0
	for true {
		for i+1 < len(A) && A[i] >= A[i+1] {
			i++
		}

		if i+1 == len(A) {
			return m
		}
		k = i

		for i+1 < len(A) && A[i] < A[i+1] {
			i++
		}

		if i+1 == len(A) {
			return m
		}

		if A[i] == A[i+1] {
			continue
		}

		for i+1 < len(A) && A[i] > A[i+1] {
			i++
		}

		if m < (i - k + 1) {
			m = i - k + 1
		}
	}

	return m
}

func main() {
	A := []int{2, 1, 4, 7, 3, 2, 5}
	// A = []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}
	// A = []int{2, 3, 3, 2, 0, 2}
	// A = []int{0, 0, 1, 0, 0, 1, 1, 1, 1, 1}
	// A = []int{875, 884, 239, 731, 723, 685}

	fmt.Println(longestMountain(A))
}
