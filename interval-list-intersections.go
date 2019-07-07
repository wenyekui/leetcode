package main

import (
	"fmt"
)

type Interval struct {
	Start int
	End   int
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intervalIntersection(A []Interval, B []Interval) []Interval {
	ans := []Interval{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		a := A[i]
		b := B[j]
		if a.Start <= b.End && b.Start <= a.End {
			start := max(a.Start, b.Start)
			end := min(a.End, b.End)
			ans = append(ans, Interval{start, end})
		}
		if a.End > b.End {
			j++
		} else {
			i++
		}
	}
	return ans
}

func main() {
	fmt.Println("")
}
