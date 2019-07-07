package main

import (
//	"fmt"
)

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	mem := make([]int, 1<<uint(maxChoosableInteger)+1)
	return dfs(0, mem, maxChoosableInteger, desiredTotal)
}

func dfs(bitmap int, mem []int, max int, target int) bool {
	if mem[bitmap] != 0 {
		return mem[bitmap] == 1
	}

	for i := 1; i <= max; i++ {
		j := 1 << uint(i-1)
		if bitmap&j != 0 {
			continue
		}
		if i >= target {
			mem[bitmap] = 1
			return true
		}

		if !dfs(bitmap|j, mem, max, target-i) {
			mem[bitmap] = 1
			return true
		}
	}
	mem[bitmap] = -1
	return false
}

func main() {
	println(canIWin(10, 11))
	println(canIWin(5, 50))
}
