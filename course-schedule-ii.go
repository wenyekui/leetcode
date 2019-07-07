package main

import (
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	g1 := make(map[int]map[int]bool)
	g2 := make(map[int][]int)

	for _, p := range prerequisites {
		a, b := p[0], p[1]
		if g1[a] == nil {
			g1[a] = map[int]bool{}
		}

		if g2[b] == nil {
			g2[b] = []int{}
		}
		g1[a][b] = true
		g2[b] = append(g2[b], a)
	}

	queue := []int{}

	for s, _ := range g2 {
		if g1[s] == nil {
			queue = append(queue, s)
		}
	}

	ans := []int{}
	for i := 0; i < numCourses; i++ {
		if g1[i] == nil && g2[i] == nil {
			ans = append(ans, i)
		}
	}

	for len(queue) != 0 {
		x := queue[0]
		ans = append(ans, x)
		queue = queue[1:]
		for _, y := range g2[x] {
			delete(g1[y], x)
			if len(g1[y]) == 0 {
				queue = append(queue, y)
			}
		}
	}

	return ans
}

func main() {
	fmt.Println("")
}
