package main

import (
	"fmt"
)

func makeMaps(A []string) []map[rune]int {
	AMaps := []map[rune]int{}
	for _, a := range A {
		amap := make(map[rune]int)
		for _, c := range a {
			amap[c]++
		}
		AMaps = append(AMaps, amap)
	}
	return AMaps
}

func wordSubsets(A []string, B []string) []string {
	AMaps := makeMaps(A)
	BMaps := makeMaps(B)
	ans := []string{}

	for i, a := range AMaps {
		flag := true
		for _, b := range BMaps {
			for l, c := range b {
				if a[l] < c {
					flag = false
					break
				}
			}
			if !flag {
				break
			}
		}
		if flag {
			ans = append(ans, A[i])
		}
	}
	return ans
}

func main() {
	fmt.Println("")
}
