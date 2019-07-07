package main

import (
	"fmt"
)

func reconstructQueue(people [][]int) [][]int {
	remain := make([]int, len(people))

	for i := 0; i < len(people); i++ {
		remain[i] = people[i][1]
	}

	for i := 0; i < len(people); i++ {

		index := len(people) - 1
		for j := i; j < len(people); j++ {
			if remain[j] == 0 && (remain[index] != 0 || (remain[index] == 0 && people[index][0] > people[j][0])) {
				index = j
			}
		}

		people[i], people[index] = people[index], people[i]
		remain[i], remain[index] = remain[index], remain[i]
		for j := i + 1; j < len(people); j++ {
			if people[j][0] <= people[i][0] {
				remain[j]--
			}
		}
	}
	return people
}

func main() {
	people := [][]int{
		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1},
		[]int{5, 0},
		[]int{6, 1},
		[]int{5, 2},
	}
	ans := reconstructQueue(people)
	fmt.Printf("%v\n", ans)
}
