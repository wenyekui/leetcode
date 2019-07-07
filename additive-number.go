package main

import (
	"fmt"
)

func add(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	res := make([]byte, len(a)+1)
	for i := 0; i < len(res); i++ {
		res[i] = '0'
	}

	var carry byte = 0

	for i := 0; i < len(a); i++ {
		var x byte

		if i >= len(b) {
			x = (a[len(a)-1-i] - '0') + carry
		} else {
			x = (b[len(b)-1-i] - '0') + (a[len(a)-1-i] - '0') + carry
		}
		if x >= 10 {
			carry = 1
			x = x % 10
		} else {
			carry = 0
		}

		res[len(res)-i-1] += x
	}
	if carry == 1 {
		res[0] += 1
	} else {
		res = res[1:]
	}
	return string(res)
}

func dfs(num string, pre []string) bool {
	if len(pre) == 2 {
		res := add(pre[0], pre[1])
		if num == res {
			return true
		}
		if len(num) < len(res) || num[:len(res)] != res {
			return false
		}
		pre[0] = pre[1]
		pre[1] = res
		return dfs(num[len(res):], pre)
	}

	if len(num) > 0 && num[0] == '0' {
		return dfs(num[1:], append(pre, "0"))
	}

	for i := 1; i < len(num); i++ {

		if dfs(num[i:], append(pre, num[:i])) {
			return true
		}
	}
	return false
}

func isAdditiveNumber(num string) bool {
	return dfs(num, []string{})
}

func main() {
	// var ans string
	var ans bool
	// ans := isAdditiveNumber("112358")
	// ans := isAdditiveNumber("199100199")
	ans = isAdditiveNumber("198019823962")
	// ans = add("1980", "1982")
	fmt.Println(ans)
}
