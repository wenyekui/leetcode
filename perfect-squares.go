package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	cache := make(map[int]int)
	return _numSquares(n, cache)
}

func _numSquares(n int, cache map[int]int) int {
	if n < 4 {
		return n
	}

	if m, ok := cache[n]; ok {
		return m
	}

	ans := n
	x := int(math.Sqrt(float64(n)))
	if x*x == n {
		return 1
	}
	for i := x; i >= 2; i-- {
		t := _numSquares(n-i*i, cache) + 1

		if t < ans {
			ans = t
		}
	}
	cache[n] = ans
	return ans
}

func main() {
	n := 12
	n = 13
	n = 144
	n = 154
	fmt.Println(numSquares(n))
}
