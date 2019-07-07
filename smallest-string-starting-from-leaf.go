package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	ans := ""
	dfs(root, []byte{}, &ans)
	return ans
}

func dfs(node *TreeNode, prev []byte, ans *string) {
	if node == nil {
		return
	} else {
		prev = append([]byte{'a' + byte(node.Val)}, prev...)
	}

	if node.Left == nil && node.Right == nil {
		if *ans == "" || string(prev) < *ans {
			*ans = string(prev)
		}
		return
	}
	dfs(node.Left, prev, ans)
	dfs(node.Right, prev, ans)
}

func main() {
	root := &TreeNode{
		0,
		&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
	}
	ans := smallestFromLeaf(root)
	fmt.Println(ans)
}
