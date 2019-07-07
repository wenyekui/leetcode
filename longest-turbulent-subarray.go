package main

import (
	"fmt"
)

func maxTurbulenceSize(A []int) int {
	if len(A) <= 1 {
		return len(A)
	}

	max := 1
	start := false
	tmp := 1
	gt := false

	for i := 0; i < len(A)-1; i++ {
		if start {
			if (gt && A[i] < A[i+1]) || (!gt && A[i] > A[i+1]) {
				gt = !gt
				tmp += 1
				if tmp > max {
					max = tmp
				}
			} else {
				start = false
				tmp = 1
			}
		}

		if !start {
			if A[i] == A[i+1] {
				continue
			}
			start = true
			tmp = 2
			if tmp > max {
				max = tmp
			}
			if A[i] > A[i+1] {
				gt = true
			} else {
				gt = false
			}
		}
	}
	return max
}

func main() {

	A := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	x := maxTurbulenceSize(A)

	A = []int{4, 8, 12, 16}
	x = maxTurbulenceSize(A)
	fmt.Println(x)
}
