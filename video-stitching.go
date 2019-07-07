package main

import (
	"fmt"
	"sort"
)

func videoStitching(clips [][]int, T int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})
	start := 0
	maxEnd := 0
	ans := 1
	for i := 0; i < len(clips); i++ {
		if clips[i][0] <= start {
			if maxEnd < clips[i][1] {
				maxEnd = clips[i][1]
			}
		} else {
			if maxEnd < clips[i][0] {
				return -1
			}
			ans++
			start = maxEnd
			maxEnd = clips[i][1]
		}
		if maxEnd >= T {
			return ans
		}

	}
	return -1
}

func main() {
	fmt.Println("")
}
