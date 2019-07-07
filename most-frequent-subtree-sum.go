package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	counter := make(map[int]int)
	max := 0
	subSum(root, counter)

	for _, v := range counter {
		if v > max {
			max = v
		}
	}

	ans := []int{}
	for k, v := range counter {
		if v == max {
			ans = append(ans, k)
		}
	}
	return ans
}

func subSum(root *TreeNode, counter map[int]int) int {
	if root == nil {
		return 0
	}
	x := root.Val + subSum(root.Left, counter) + subSum(root.Right, counter)
	counter[x] += 1
	return x
}

func main() {
	fmt.Println("")
}
