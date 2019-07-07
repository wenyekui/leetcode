package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := bstToGst(root.Right)
	if node != nil {
		for node.Left != nil {
			node = node.Left
		}
		root.Val += right.Val
	}
	if root.Left != nil {
		node = root.Left
		for node.Right != nil {
			node = node.Right
		}
		node.Val += root.Val
	}
	bstToGst(root.Left)
	return root
}

func main() {
	fmt.Println("")
}
