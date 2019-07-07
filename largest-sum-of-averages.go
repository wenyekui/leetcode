package main

import (
	"fmt"
)

func largestSumOfAverages(A []int, K int) float64 {

	if K == 0 {
		return 0
	}
	var max float64 = 0
	for i := 0; i+K-1 < len(A); i++ {
		t1 := avg(A[0 : i+1])
		t2 := largestSumOfAverages(A[i+1:len(A)], K-1)
		if t1+t2 > max {
			max = t1 + t2
		}
	}

	fmt.Println(A)
	fmt.Println(K)
	println(max)
	return max
}

func sum(A []int) float64 {
	s := 0
	for _, a := range A {
		s += a
	}
	return float64(s)
}

func avg(A []int) float64 {
	return sum(A) / float64(len(A))
}

func main() {
	//A := []int{9, 1, 2, 3, 9}
	//K := 3

	A := []int{9}
	K := 1

	A = []int{4, 1, 7, 5, 6, 2, 3}
	K = 4
	fmt.Println(largestSumOfAverages(A, K))
}
