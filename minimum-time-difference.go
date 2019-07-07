package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinDifference(timePoints []string) int {
	N := len(timePoints)
	minutes := make([]int, N)
	for i := 0; i < N; i++ {
		minutes[i] = parseTime(timePoints[i])
	}
	sort.Ints(minutes)

	ans := 60 * 24

	for i := 1; i < N; i++ {
		ans = min(ans, minutes[i]-minutes[i-1])
	}

	ans = min(ans, minutes[0]+60*24-minutes[N-1])
	return ans
}

func parseTime(time string) int {
	parts := strings.Split(time, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return h*60 + m
}

func main() {
	fmt.Println("")
}
