package main

import (
	"fmt"
)

func findClosestElements(arr []int, k int, x int) []int {
	lo := 0
	hi := len(arr)
	for lo < hi {
		m := (lo + hi) / 2
		if x > arr[m] {
			lo = m + 1
		} else {
			hi = m
		}
	}
}

func main() {
	fmt.Println("")
}
