package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
}

type SortedInterval struct {
	Start, End, Index int
}
type IntervalsSort []SortedInterval

func (a IntervalsSort) Len() int           { return len(a) }
func (a IntervalsSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntervalsSort) Less(i, j int) bool { return a[i].Start < a[j].Start }

func findRightInterval(intervals []Interval) []int {
	sorted := make([]SortedInterval, len(intervals))
	for i, interval := range intervals {
		sorted[i] = SortedInterval{interval.Start, interval.End, i}
	}

	sort.Sort(IntervalsSort(sorted))

	binarySearch := func(value int) SortedInterval {
		start := 0
		end := len(sorted) - 1
		for start <= end {
			mid := (start + end) / 2
			interval := sorted[mid]

			if interval.Start == value {
				return interval
			}

			if interval.Start > value {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

		return sorted[start]
	}

	last := sorted[len(sorted)-1].Start

	result := make([]int, len(sorted))
	for _, interval := range sorted {
		if interval.End > last {
			result[interval.Index] = -1
			continue
		}
		result[interval.Index] = binarySearch(interval.End).Index
	}

	//fmt.Println(intervals)
	//fmt.Println(sorted)

	return result
}
