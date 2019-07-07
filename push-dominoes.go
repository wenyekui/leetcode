package main

import (
	"fmt"
)

func pushDominoes(dominoes string) string {

	LList := []int{}
	RList := []int{}

	ans := []byte(dominoes)

	for i := 0; i < len(dominoes); i++ {
		switch dominoes[i] {
		case 'R':
			RList = append(RList, i)
		case 'L':
			LList = append(LList, i)
		}
	}

	for len(LList) != 0 || len(RList) != 0 {
		keep := map[int]bool{}
		tmp := []int{}

		for _, i := range LList {
			if i-1 >= 0 && ans[i-1] == '.' {
				if i-2 >= 0 && ans[i-2] == 'R' {
					keep[i-1] = true
				} else {
					ans[i-1] = 'L'
					tmp = append(tmp, i-1)
				}
			}
		}
		LList = tmp
		tmp = []int{}

		for _, i := range RList {
			if i+1 < len(ans) && ans[i+1] == '.' && !keep[i+1] {
				ans[i+1] = 'R'
				tmp = append(tmp, i+1)
			}
		}
		RList = tmp
	}

	return string(ans)
}

func main() {
	fmt.Println("")
}
