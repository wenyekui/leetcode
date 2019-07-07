package main

import (
	"fmt"
)

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, asteroid := range asteroids {
		for len(stack) != 0 && asteroid < 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -asteroid {
			stack = stack[0 : len(stack)-1]
		}
		if len(stack) == 0 || stack[len(stack)-1]*asteroid > 0 || asteroid > 0 {
			stack = append(stack, asteroid)
			continue
		}
		if stack[len(stack)-1] <= -asteroid {
			stack = stack[0 : len(stack)-1]
		}
	}
	return stack
}

func main() {
	asteroids := []int{5, 10, -5}
	asteroids = []int{-2, -1, 1, 2}
	fmt.Println(asteroidCollision(asteroids))
}
