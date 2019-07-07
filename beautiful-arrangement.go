package main

import (
	"fmt"
)

func countArrangement(N int) int {
	a := make([]int, N, N)
	b := make([]int, N, N)
	for i := 1; i <= N; i++ {
		a[i-1] = i
		b[i-1] = i
	}
	return dfs(a, b)
}

func dfs(a []int, b []int) int {
	n := len(a)
	if n == 0 {
		return 1
	}
	result := 0
	for i := 0; i < n; i++ {
		a[i], a[0] = a[0], a[i]
		if a[0]%b[0] == 0 || b[0]%a[0] == 0 {
			result += dfs(a[1:n], b[1:n])
		}
		a[i], a[0] = a[0], a[i]
	}
	return result
}

//var count int
//
//func countArrangement(N int) int {
//	count = 0
//	used := make([]bool, N+1, N+1)
//	dfs(N, used, 1)
//	return count
//}
//
//func dfs(N int, used []bool, pos int) {
//	if pos > N {
//		count++
//	}
//	for i := 1; i <= N; i++ {
//		if !used[i] && (i%pos == 0 || pos%i == 0) {
//			used[i] = true
//			dfs(N, used, pos+1)
//			used[i] = false
//		}
//	}
//}

func main() {
	fmt.Println(countArrangement(2))
}
