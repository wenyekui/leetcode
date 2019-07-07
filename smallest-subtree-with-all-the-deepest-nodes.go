package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	ans, _ := dfs(root, 0)
	return ans
}

func dfs(node *TreeNode, depth int) (*TreeNode, int) {
	if node == nil {
		return node, depth
	}
	ln, ld := dfs(node.Left, depth+1)
	rn, rd := dfs(node.Right, depth+1)

	if ld == rd {
		return node, ld
	}
	if ld > rd {
		return ln, ld
	}
	return rn, rd
}

func main() {
	fmt.Println("")
}
