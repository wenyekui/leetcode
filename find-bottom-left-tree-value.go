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

func findBottomLeftValue(root *TreeNode) (ans int) {
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		ans = nodes[0].Val
		tmp := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}

		}
		nodes = tmp
	}
	return
}

func main() {
	fmt.Println("")
}
