package main

import (
	"fmt"
)

func networkDelayTime(times [][]int, N int, K int) int {
	graph := make(map[int][][]int)
	for _, time := range times {
		if graph[time[0]] == nil {
			graph[time[0]] = make([][]int, 0, N)
		}
		graph[time[0]] = append(graph[time[0]], []int{time[1], time[2]})
	}
	visited := make([]bool, N+1, N+1)
	ans := bfs(graph, K, visited)
	for i := 1; i <= N; i++ {
		if !visited[i] {
			return -1
		}
	}
	return ans
}

func bfs(graph map[int][][]int, K int, visited []bool) int {
	ans := 0
	queue := [][]int{[]int{K, 0}}
	for len(queue) != 0 {
		min_i := 0
		for i := 1; i < len(queue); i++ {
			if queue[min_i][1] > queue[i][1] {
				min_i = i
			}
		}
		queue[0], queue[min_i] = queue[min_i], queue[0]
		x := queue[0]
		queue = queue[1:]
		if visited[x[0]] {
			continue
		}
		visited[x[0]] = true
		for i := 0; i < len(queue); i++ {
			queue[i][1] -= x[1]
		}
		ans += x[1]
		if graph[x[0]] == nil {
			continue
		}
		for _, next := range graph[x[0]] {
			if visited[next[0]] {
				continue
			}
			queue = append(queue, next)
		}
	}
	return ans
}

func main() {
	times := [][]int{
		[]int{3, 1, 1},
		[]int{2, 3, 1},
		[]int{3, 4, 1},
	}
	N := 4
	K := 2

	times = [][]int{
		[]int{1, 2, 1},
		[]int{2, 1, 3},
	}
	N = 2
	K = 2

	times = [][]int{
		[]int{1, 2, 1},
		[]int{2, 3, 2},
		[]int{1, 3, 2},
	}
	N = 3
	K = 1
	ans := networkDelayTime(times, N, K)
	fmt.Println(ans)
}
