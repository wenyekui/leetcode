package main

import (
	"fmt"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}

	counter := make(map[int]map[int]bool)

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if counter[x] == nil {
			counter[x] = map[int]bool{}
		}
		if counter[y] == nil {
			counter[y] = map[int]bool{}
		}
		counter[x][y] = true
		counter[y][x] = true
	}

	for len(counter) > 2 {
		tmp := []int{}
		for k, m := range counter {
			if len(m) == 1 {
				tmp = append(tmp, k)
			}
		}
		for _, k := range tmp {
			tmp2 := []int{}
			for k2, _ := range counter[k] {
				tmp2 = append(tmp2, k2)
			}
			delete(counter[tmp2[0]], k)
			delete(counter, k)
		}
	}

	fmt.Printf("%v\n", counter)

	ans := []int{}
	for k, _ := range counter {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	edges := [][]int{
		// []int{0, 3},
		// []int{1, 3},
		// []int{2, 3},
		// []int{4, 3},
		// []int{5, 4},

		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{3, 4},
		[]int{4, 5},
	}
	n := 6
	ans := findMinHeightTrees(n, edges)
	fmt.Println(ans)
}
