package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	digits := []byte(num)
	ans := []byte{}
	removed := make(map[int]bool)
	for i := 0; i < len(digits) && k != 0; i++ {
		if i == len(digits)-1 || digits[i] > digits[i+1] {
			for j := i; j >= 0 && k != 0; j-- {
				if i != len(digits)-1 && digits[j] <= digits[i+1] {
					break
				}

				if removed[j] {
					continue
				}

				k--
				removed[j] = true
			}
			continue
		}
	}
	for i, _ := range digits {
		if removed[i] {
			continue
		}
		if len(ans) != 0 || digits[i] != '0' {
			ans = append(ans, digits[i])
		}
	}
	if len(ans) == 0 {
		return "0"
	}
	return string(ans)
}

func main() {
	ans := removeKdigits("10", 2)
	fmt.Println(ans)
	ans = removeKdigits("1432219", 3)
	// ans = removeKdigits("10", 2)
	fmt.Println(ans)

	ans = removeKdigits("10", 1)
	fmt.Println(ans)
}
