package main

import (
	"fmt"
)

func minAddToMakeValid(S string) int {
	stack := make([]rune, 0, 1000)
	for _, b := range S {
		if len(stack) == 0 || b == '(' || stack[len(stack)-1] == ')' {
			stack = append(stack, b)
			continue
		}
		stack = stack[0 : len(stack)-1]
	}
	return len(stack)
}

func main() {
	println(minAddToMakeValid("())"))
	println(minAddToMakeValid("((("))
	println(minAddToMakeValid("())"))

	fmt.Println("")
}
