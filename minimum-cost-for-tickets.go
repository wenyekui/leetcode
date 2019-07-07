package main

import (
	"fmt"
)

func min(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if m > num {
			m = num
		}
	}
	return m
}

func mincostTickets(days []int, costs []int) int {
	N := len(days)
	dp := make([]int, N)
	dp[N-1] = min(costs...)

	for i := N - 2; i >= 0; i-- {
		j := i + 1
		c1 := costs[0] + dp[j]
		c2 := costs[1]
		if days[i]+7 <= days[N-1] {
			for days[j] < days[i]+7 {
				j++
			}
			c2 += dp[j]
		}

		c3 := costs[2]
		if days[i]+30 <= days[N-1] {
			for days[j] < days[i]+30 {
				j++
			}
			c3 += dp[j]
		}
		dp[i] = min(c1, c2, c3)
	}

	return dp[0]
}

func main() {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	ans := mincostTickets(days, costs)
	fmt.Println(ans)
	days = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs = []int{2, 7, 15}
	ans = mincostTickets(days, costs)
	fmt.Println(ans)
}
