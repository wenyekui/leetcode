package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return 0
	}

	i, j, r := 0, 0, 0

	for true {
		r = r + gas[j] - cost[j]
		for r < 0 {

			if i == 0 {
				i = len(gas) - 1
			} else {
				i--
			}
			if i == j {
				return -1
			}
			r = r + gas[i] - cost[i]
		}
		if (j+1)%len(gas) == i {
			break
		}
		j++
	}

	return i
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	ans := canCompleteCircuit(gas, cost)

	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	ans = canCompleteCircuit(gas, cost)
	fmt.Println(ans)
}
