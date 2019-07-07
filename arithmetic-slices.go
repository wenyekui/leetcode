package main

import (
	"fmt"
)

func numberOfArithmeticSlices(A []int) int {
	N := len(A)
	if N < 3 {
		return 0
	}
	ans := 0
	n := 2
	d := A[1] - A[0]
	for i := 2; i < N; i++ {
		if A[i]-A[i-1] == d {
			n++
		} else {
			ans += (n - 2) * (n - 1) / 2
			d = A[i] - A[i-1]
			n = 2
		}
	}

	ans += (n - 2) * (n - 1) / 2
	return ans
}

func main() {
	A := []int{1, 2, 3, 4}
	fmt.Println(numberOfArithmeticSlices(A))
}
