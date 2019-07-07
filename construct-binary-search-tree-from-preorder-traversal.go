package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var i int
	for i = 1; i < len(preorder); i++ {
		if preorder[i] > preorder[0] {
			break
		}
	}

	leftValues := preorder[1:i]
	rightValues := preorder[i:len(preorder)]

	root := &TreeNode{preorder[0], bstFromPreorder(leftValues), bstFromPreorder(rightValues)}
	return root
}

func main() {
	fmt.Println("")
}
