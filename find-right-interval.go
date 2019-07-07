package main

import (
	"fmt"
	//	"sort"
)

/**
 * Definition for an interval.
 */

type Interval struct {
	Start int
	End   int
}

func findRightInterval(intervals []Interval) []int {
	N := len(intervals)
	ans := []int{}
	for i := 0; i < N; i++ {
		var j int
		for j = 0; j < N; j++ {
			if intervals[i].End == intervals[j].Start {
				ans = append(ans, j)
				break
			}
		}

		if j == N {
			for j = 0; j < N; j++ {
				if intervals[i].End < intervals[j].Start {
					ans = append(ans, j)
					break
				}
			}
			if j == N {
				ans = append(ans, -1)
			}
		}
	}
	return ans
}

func main() {
	intervals := []Interval{Interval{1, 2}}

	intervals = []Interval{Interval{3, 4}, Interval{2, 3}, Interval{1, 2}}
	fmt.Println(findRightInterval(intervals))
}
