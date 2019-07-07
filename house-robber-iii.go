package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans1 := root.Val
	if root.Left != nil {
		ans1 += rob(root.Left.Left)
		ans1 += rob(root.Left.Right)
	}
	if root.Right != nil {
		ans1 += rob(root.Right.Left)
		ans1 += rob(root.Right.Right)
	}

	ans2 := rob(root.Left) + rob(root.Right)
	if ans1 > ans2 {
		return ans1
	}
	return ans2
}

func main() {
	fmt.Println("")
}
