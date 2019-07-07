package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type NodeWrapper struct {
	*TreeNode
	depth  int
	offset int
}

func widthOfBinaryTree(root *TreeNode) int {
	ans := 0
	queue := []*NodeWrapper{&NodeWrapper{root, 1, 1}}
	curDepth := 0
	beginOffset := 0
	for len(queue) != 0 {
		node := queue[0]
		if node.depth != curDepth {
			curDepth = node.depth
			beginOffset = node.offset
		}

		if node.offset-beginOffset+1 > ans {
			ans = node.offset - beginOffset + 1
		}

		if node.Left != nil {
			queue = append(queue, &NodeWrapper{node.Left, node.depth + 1, node.offset*2 - 1})
		}

		if node.Right != nil {
			queue = append(queue, &NodeWrapper{node.Right, node.depth + 1, node.offset * 2})
		}
		queue = queue[1:len(queue)]
	}

	return ans
}

func main() {
	fmt.Println("")
}
