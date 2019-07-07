package main

import (
	"fmt"
	"strings"
)

func numJewelsInStones(J string, S string) int {
	count := 0
	for _, s := range S {
		if strings.ContainsRune(J, s) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("")
}
