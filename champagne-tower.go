package main

import (
	"fmt"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	a := (1 + query_row) * query_row / 2
	println(a)
	if poured <= a {
		return 0
	}
	b := (1 + query_row + 1) * (query_row + 1) / 2
	println(b)
	if poured >= b {
		return 1
	}
	var c float64 = float64(poured-a) / float64(query_row)
	println(c)
	if query_glass != 0 && query_glass != query_row {
		return c
	}
	return c / 2

}

func main() {
	poured := 1
	query_glass := 1
	query_row := 1

	poured = 2
	query_glass = 1
	query_row = 1

	poured = 6
	query_row = 2
	query_glass = 0

	ans := champagneTower(poured, query_row, query_glass)

	fmt.Println(ans)
}
