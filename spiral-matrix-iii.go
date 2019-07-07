package main

import (
	"fmt"
)

func add(ans [][]int, R int, C int, r int, c int) [][]int {
	if r >= 0 && r < R && c >= 0 && c < C {
		ans = append(ans, []int{r, c})
	}
	return ans
}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	ans := make([][]int, 0, R*C)

	d := 1
	r := r0
	c := c0
	ans = append(ans, []int{r0, c0})

	for true {
		if d > 2*R && d > 2*C {
			break
		}

		for i := 0; i < d; i++ {
			c += 1
			ans = add(ans, R, C, r, c)
		}
		for i := 0; i < d; i++ {
			r += 1
			ans = add(ans, R, C, r, c)
		}

		d += 1
		for i := 0; i < d; i++ {
			c -= 1
			ans = add(ans, R, C, r, c)
		}

		for i := 0; i < d; i++ {
			r -= 1
			ans = add(ans, R, C, r, c)
		}
		d += 1
	}
	return ans
}

func main() {
	x := spiralMatrixIII(1, 4, 0, 0)
	fmt.Println(x)

	x = spiralMatrixIII(5, 6, 1, 4)
	fmt.Println(x)
}
