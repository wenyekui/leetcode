package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	ans := 0
	sort.Ints(people)
	x := 0
	for i := len(people) - 1; i >= 0; i-- {
		if i < x {
			break
		}
		if people[i] == limit {
			ans += 1
			continue
		}

		if i == x {
			ans += 1
			break
		}

		if people[i]+people[x] <= limit {
			ans += 1
			x += 1
		} else {
			ans += 1
		}
	}
	return ans
}

func main() {
	people := []int{1, 2}
	limit := 3

	x := numRescueBoats(people, limit)
	fmt.Println(x)
	people = []int{3, 5, 3, 4}
	limit = 5
	x = numRescueBoats(people, limit)
	fmt.Println(x)
}
