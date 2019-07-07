package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func dfs(node *TreeNode, max int, min int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return dfs(node.Left, node.Val, min) && dfs(node.Right, max, node.Val)
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, MaxInt, MinInt)
}

func main() {
	fmt.Println("")
}
