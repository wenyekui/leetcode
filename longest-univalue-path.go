package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {

}

func do(node *TreeNode, parentVal int, curLen int, maxLen int) int {
	if node == nil {
		return maxLen
	}
	if node.Val == parentVal {

	}
}

func main() {
	fmt.Println("")
}
