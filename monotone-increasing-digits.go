package main

import (
	"fmt"
)

func monotoneIncreasingDigits(N int) int {
	digits := toDigits(N)
	i := len(digits) - 1
	for i >= 1 {
		if digits[i] <= digits[i-1] {
			i--
			continue
		}

		for j := i - 1; j >= 0; j-- {
			digits[j] = 9
		}
		digits[i] -= 1
		for i < len(digits)-1 {
			if digits[i] < digits[i+1] {
				digits[i] = 9
				digits[i+1] -= 1
			} else {
				break
			}
			i++
		}
		break
	}
	return toNumber(digits)
}

func toNumber(digits []int) int {
	res := 0
	for i := len(digits) - 1; i >= 0; i-- {
		res = res*10 + digits[i]
	}
	return res
}

func toDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := []int{}
	for n != 0 {
		x := n % 10
		res = append(res, x)
		n = n / 10
	}
	return res
}

func main() {
	x := monotoneIncreasingDigits(10)
	x = monotoneIncreasingDigits(1234)
	x = monotoneIncreasingDigits(332)
	fmt.Println(x)
}
