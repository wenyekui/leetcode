package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	N := len(A)
	maxLArray := make([]int, N)
	maxMArray := make([]int, N)
	lSumArray := make([]int, N)
	rSumArray := make([]int, N)

	copy(lSumArray, A)
	copy(rSumArray, A)

	for i := 1; i < N; i++ {
		lSumArray[i] += lSumArray[i-1]
		maxLArray[i] = lSumArray[i]
		maxMArray[i] = lSumArray[i]
	}
	for i := N - 2; i >= 0; i-- {
		rSumArray[i] += rSumArray[i+1]
	}

	maxL := lSumArray[L-1]
	maxM := lSumArray[M-1]

	startL := 0
	startM := 0

	for i := L; i < len(A); i++ {
		maxL = max(maxL, lSumArray[i]-lSumArray[startL])
		maxLArray[i] = maxL
		startL++
	}

	for i := M; i < len(A); i++ {
		maxM = max(maxM, lSumArray[i]-lSumArray[startM])
		maxMArray[i] = maxM
		startM++
	}

	startR := N - 1
	maxR := rSumArray[N-M]
	ans := maxR + maxLArray[N-M-1]
	for i := N - M - 1; i >= L; i-- {
		maxR = max(maxR, rSumArray[i]-rSumArray[startR])
		ans = max(ans, maxR+maxLArray[i-1])
		startR--
	}
	startR = N - 1
	maxR = rSumArray[N-L]
	ans = max(ans, maxR+maxMArray[N-L-1])
	for i := N - L - 1; i >= M; i-- {
		maxR = max(maxR, rSumArray[i]-rSumArray[startR])
		ans = max(ans, maxR+maxMArray[i-1])
		startR--
	}
	return ans
}

func main() {
	A := []int{0, 6, 5, 2, 2, 5, 1, 9, 4}
	L := 1
	M := 2

	A = []int{3, 8, 1, 3, 2, 1, 8, 9, 0}
	L = 3
	M = 2

	A = []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}
	L = 4
	M = 3
	ans := maxSumTwoNoOverlap(A, L, M)

	fmt.Println(ans)
}
