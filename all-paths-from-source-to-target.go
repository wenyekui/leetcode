package main

import (
	"fmt"
)

var result [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	result = make([][]int, 0)
	dfs(graph, []int{0})
	return result
}

func dfs(graph [][]int, cur []int) {
	i := cur[len(cur)-1]
	if i == len(graph)-1 {
		c := make([]int, len(cur))
		copy(c, cur)
		result = append(result, c)
		return
	}
	m := len(graph[i])
	for j := 0; j < m; j++ {
		dfs(graph, append(cur, graph[i][j]))
	}
}

func main() {
	graph := [][]int{
		[]int{1, 2},
		[]int{3},
		[]int{3},
		[]int{},
	}
	result = allPathsSourceTarget(graph)
	fmt.Println(result)
}
