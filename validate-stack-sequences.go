package main

import (
// 	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	if len(popped) > len(pushed) {
		return false
	}

	stack := []int{}
	i := 0
	for _, val := range popped {
		if len(stack) == 0 || stack[len(stack)-1] != val {
			for i < len(pushed) && pushed[i] != val {
				stack = append(stack, pushed[i])
				i++
			}
			if i == len(pushed) {
				return false
			}
			i++
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}
	return true
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}

	pushed = []int{1, 2, 3, 4, 5}
	popped = []int{4, 3, 5, 1, 2}
	println(validateStackSequences(pushed, popped))
}
