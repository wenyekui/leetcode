package main

import (
	"fmt"
)

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func countSubstrings(s string) int {
	ss := make([]rune, 0, 2*len(s))
	ss = append(ss, '#')
	for _, c := range s {
		ss = append(ss, c)
		ss = append(ss, '#')
	}

	maxRight := 0
	pos := 0
	P := make([]int, len(ss))
	for i := 1; i < len(ss); i++ {
		left := 2*pos - maxRight
		r := 0
		if i < maxRight {
			j := 2*pos - i
			if j-P[j] > left {
				P[i] = P[j]
				continue
			} else {
				r = P[j]
			}
		}

		r++
		for i-r >= 0 && i+r < len(ss) {
			if ss[i-r] == ss[i+r] {
				r++
			} else {
				break
			}
		}
		r--
		P[i] = r
		if i+r > maxRight {
			maxRight = i + r
			pos = i
		}
	}
	ans := 0
	for i := 0; i < len(P); i++ {
		if P[i]%2 != 0 {
			P[i] += 1
		}
		ans += P[i] / 2
	}
	return ans
}

func main() {
	s := "abc"
	// s = "aaa"
	ans := countSubstrings(s)
	fmt.Println(ans)
}
