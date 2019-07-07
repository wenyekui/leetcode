package main

import (
	"fmt"
)

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

func toNumber(digits []int, x int) int {
	res := 0
	for i := len(digits) - 1; i > x; i-- {
		res = res*10 + digits[i]
	}
	for i := 0; i <= x; i++ {
		res = res*10 + digits[i]
	}
	return res
}

func nextGreaterElement(n int) int {
	digits := toDigits(n)

	i := 0

	for i < len(digits)-1 {
		if digits[i] > digits[i+1] {
			break
		}
		i++
	}

	if i == len(digits)-1 {
		return -1
	}

	for j := 0; j < len(digits); j++ {
		if digits[j] > digits[i+1] {
			digits[j], digits[i+1] = digits[i+1], digits[j]
			break
		}
	}
	ans := toNumber(digits, i)
	if ans > 2147483647 {
		return -1
	}
	return ans
}

func main() {
	n := 12
	n = 21
	n = 230241
	ans := nextGreaterElement(n)
	fmt.Println(ans)
}
