package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			for len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]
				if node != nil {
					return false
				}
			}
			return true
		}
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return true
}

func main() {

	fmt.Println("")
}
