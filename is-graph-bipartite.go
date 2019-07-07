package main

import (
	"fmt"
)

func isBipartite(graph [][]int) bool {
	groups := make([][]bool, 2)
	for i := 0; i < len(groups); i++ {
		groups[i] = make([]bool, len(graph))
	}
}

func dfs(graph [][]int, i int, groups [][]bool, gid int) bool {
	groups[gid][i] = true
	gid2 := (gid + 1) % 2
	for _, j := range graph[i] {
		if groups[gid2][j] {
			continue
		}
		if groups[gid][j] {
			return false
		}
		groups[gid2][j] = true
		if !dfs(graph, j, groups, gid2) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("")
}
