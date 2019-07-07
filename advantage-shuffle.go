package main

import (
	"fmt"
	"sort"
)

func advantageCount(A []int, B []int) []int {
	N := len(A)
	ans := make([]int, N)
	indices := make([]int, N)
	for i := 0; i < N; i++ {
		indices[i] = i
	}
	sort.Ints(A)
	sort.Slice(indices, func(i, j int) bool {
		return B[indices[i]] < B[indices[j]]
	})

	start := 0
	end := N - 1

	fmt.Println("indices", indices)

	for i := 0; i < N; i++ {
		fmt.Println("ans", ans)
		if A[i] > B[indices[start]] {
			ans[indices[start]] = A[i]
			start++
		} else {
			ans[indices[end]] = A[i]
			end--
		}
	}
	return ans
}

func main() {

	A := []int{5621, 1743, 5532, 3549, 9581}
	B := []int{913, 9787, 4121, 5039, 1481}

	fmt.Println("A", A)
	fmt.Println("B", B)

	ans := advantageCount(A, B)

	sort.Ints(A)
	sort.Ints(B)

	fmt.Println("sorted A", A)
	fmt.Println("sorted B", B)

	fmt.Println(ans)
}
