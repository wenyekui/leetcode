package main

import (
	"fmt"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)

	for i, equation := range equations {
		x, y := equation[0], equation[1]

		if graph[x] == nil {
			graph[x] = make(map[string]float64)
		}

		if graph[y] == nil {
			graph[y] = make(map[string]float64)
		}
		graph[x][y] = values[i]
		graph[y][x] = 1 / values[i]
	}
	ans := make([]float64, len(queries))
	for i, query := range queries {
		ans[i] = dfs(graph, query[0], query[1], 1, map[string]bool{})
	}
	return ans
}

func dfs(graph map[string]map[string]float64, x, y string, res float64, prev map[string]bool) float64 {
	if graph[x] == nil {
		return -1
	}
	if x == y {
		return res
	}
	for k, v := range graph[x] {
		if prev[k] {
			continue
		}
		prev[k] = true
		ans := dfs(graph, k, y, res*v, prev)
		prev[k] = false
		if ans != -1 {
			return ans
		}
	}
	return -1
}

func main() {
	fmt.Println("")
}
