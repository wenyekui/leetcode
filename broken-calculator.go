package main

import (
	"fmt"
)

func brokenCalc(X int, Y int) int {
	if X == Y {
		return 0
	}
	if X > Y {
		return X - Y
	}
	if Y%2 == 0 {
		return brokenCalc(X, Y/2) + 1
	} else {
		return brokenCalc(X, Y+1) + 1
	}
}

func main() {

	x := brokenCalc(2, 3)
	x = brokenCalc(1024, 1)
	fmt.Println(x)
}
