package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	sum int
	i   int
	j   int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].sum < pq[j].sum
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func leastBricks(wall [][]int) int {
	if len(wall) == 0 || len(wall[0]) == 0 {
		return 0
	}

	width := 0
	for _, w := range wall[0] {
		width += w
	}
	pq := make(PriorityQueue, 0, len(wall)*len(wall[0]))
	heap.Init(&pq)

	counter := make(map[int]int)
	max := 0

	for i, _ := range wall {
		heap.Push(&pq, &Item{0, i, 0})
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.j < len(wall[item.i]) {
			counter[item.sum]--
			item.sum += wall[item.i][item.j]
			counter[item.sum]++
			item.j++
			if item.sum < width && counter[item.sum] > max {
				max = counter[item.sum]
			}
			heap.Push(&pq, item)
		}

	}
	return len(wall) - max
}

func main() {

	wall := [][]int{
		[]int{1, 1},
		[]int{2},
		[]int{1, 1},
	}
	ans := leastBricks(wall)
	fmt.Println(ans)
}
