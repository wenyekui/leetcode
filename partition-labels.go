package main

import (
	"fmt"
)

func _partitionLabels(S string) []int {
	partitions := make([]int, 0, len(S))
	m := make(map[byte]int)

	for i, c := range []byte(S) {
		if partition, ok := m[c]; !ok {
			partitions = append(partitions, i)
			m[c] = len(partitions) - 1
		} else {
			for k := partitions[partition]; k <= i; k++ {
				m[S[k]] = partition
			}
			partitions = partitions[0 : partition+1]
			partitions[partition] = i

		}
	}
	result := make([]int, 0, len(partitions))
	start := 0
	for _, num := range partitions {
		result = append(result, num-start+1)
		start = num + 1
	}
	return result
}

func partitionLabels(S string) []int {
	var m [26]int
	for i, c := range []byte(S) {
		m[c-'a'] = i
	}
	list := make([]int, 0)
	start := 0
	last := 0
	for i, c := range []byte(S) {
		last = max(last, m[c-'a'])
		if last == i {
			list = append(list, last-start+1)
			start = last + 1
		}
	}

	return list
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	S := "ababcbacadefegdehijhklij"
	r := partitionLabels(S)
	fmt.Println(r)
}
