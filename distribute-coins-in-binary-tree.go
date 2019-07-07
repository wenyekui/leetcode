package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	ans := 0
	x(root, &ans)
	return ans
}

func x(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	l := x(root.Left, ans)
	r := x(root.Right, ans)
	res := root.Val + l + r - 1
	if res < 0 {
		*ans -= res
	} else {
		*ans += res
	}
	return res
}

func main() {
	fmt.Println("")
}
