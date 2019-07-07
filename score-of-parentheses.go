package main

import (
	"fmt"
)

func scoreOfParentheses(S string) int {
	stack := []int{0}

	for _, s := range S {
		if s == '(' {
			stack = append(stack, -1)
			continue
		}
		score := 1
		if stack[len(stack)-1] < 0 {
			stack = stack[0 : len(stack)-1]
		} else {
			score = 2 * stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
		}

		if stack[len(stack)-1] < 0 {
			stack = append(stack, score)
		} else {
			stack[len(stack)-1] += score
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(scoreOfParentheses("()"))
	fmt.Println(scoreOfParentheses("(())"))
	fmt.Println(scoreOfParentheses("()()"))
	fmt.Println(scoreOfParentheses("(()(()))"))
}
