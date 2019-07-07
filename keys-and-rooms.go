package main

import (
	"fmt"
)

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return false
	}
	visited := make([]bool, len(rooms), len(rooms))
	dfs(visited, rooms, 0)
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func dfs(visited []bool, rooms [][]int, current int) {
	if visited[current] {
		return
	}
	visited[current] = true

	for _, key := range rooms[current] {
		dfs(visited, rooms, key)
	}
}

func main() {
	fmt.Println("")
}
