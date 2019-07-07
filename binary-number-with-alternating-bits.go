package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	prev := n % 2
	n = n / 2
	for n != 0 {
		cur := n % 2
		if prev == cur {
			return false
		}
		prev = cur
		n = n / 2
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
}
