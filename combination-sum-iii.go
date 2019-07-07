package main

import (
	"fmt"
)

func combinationSum3(k int, n int) [][]int {
	return dfs(k, n, 1)
}

func dfs(k int, n int, i int) [][]int {

	if k == 0 && n == 0 {
		return [][]int{[]int{}}
	}
	if k <= 0 {
		return [][]int{}
	}
	ans := [][]int{}
	for j := i; j <= 9; j++ {
		if n < j {
			continue
		}
		res := dfs(k-1, n-j, j+1)
		if len(res) == 0 {
			continue
		}
		for _, r := range res {
			r = append(r, j)
			ans = append(ans, r)
		}
	}
	return ans
}

func main() {
	ans := combinationSum3(3, 7)
	ans = combinationSum3(3, 9)
	fmt.Printf("%v\n", ans)
}
