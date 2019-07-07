package main

import (
	"fmt"
)

func fourSumCount(A []int, B []int, C []int, D []int) int {
	preSum := make(map[int]int, len(A))

	for _, a := range A {
		for _, b := range B {
			preSum[a+b] += 1
		}
	}
	res := 0
	for _, c := range C {
		for _, d := range D {
			res += preSum[-c-d]
		}
	}
	return res
}

// func fourSumCount(A []int, B []int, C []int, D []int) int {
// 	res := 0
// 	for _, a := range A {
// 		for _, b := range B {
// 			for _, c := range C {
// 				for _, d := range D {
// 					if a+b+c+d == 0 {
// 						res++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
