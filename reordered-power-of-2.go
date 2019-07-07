package main

import (
	"fmt"
)

func reorderedPowerOf2(N int) bool {
	if N == 1 {
		return true
	}
	c := countDigits(N)
	x := 1
	for i := 1; i < 30; i++ {
		x = 2 * x
		m := countDigits(x)
		flag := true
		for j := 0; j <= 9; j++ {
			if c[j] != m[j] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func countDigits(n int) map[int]int {
	res := make(map[int]int)
	for n != 0 {
		d := n % 10
		res[d] += 1
		n = n / 10
	}
	return res
}

func main() {
	N := 46
	x := reorderedPowerOf2(N)
	N = 24
	x = reorderedPowerOf2(N)

	N = 10
	x = reorderedPowerOf2(N)
	fmt.Println(x)
}
