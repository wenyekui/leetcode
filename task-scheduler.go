package main

import (
	"fmt"
)

func leastInterval(tasks []byte, n int) int {
	counter := map[byte]int{}
	max := 0
	for _, t := range tasks {
		counter[t] += 1
		if max < counter[t] {
			max = counter[t]
		}
	}
	maxCount := 0
	for _, c := range counter {
		if c == max {
			maxCount++
		}
	}
	return (max-1)*(n+1) + maxCount
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(tasks, n))
}
