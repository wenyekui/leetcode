package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{val, nil, nil}
	if root == nil {
		return node
	}

	if val < root.Val {
		if root.Left == nil {
			root.Left = node
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root
}
