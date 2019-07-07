package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans, _ := dfs(root, p, q)
	return ans
}

func dfs(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}

	node1, b1 := dfs(root.Left, p, q)
	if node1 != nil {
		return node1, true
	}

	node2, b2 := dfs(root.Right, p, q)
	if node2 != nil {
		return node2, true
	}
	if b1 && b2 {
		return root, true
	}

	got := false
	var anc *TreeNode
	if root == p || root == q {
		got = true
	}

	if b1 || b2 {
		if got {
			anc = root
		}
		got = true
	}
	return anc, got
}

func main() {
	root := &TreeNode{
		3,
		&TreeNode{
			5, &TreeNode{
				6, nil, nil,
			}, &TreeNode{
				2, &TreeNode{
					7, nil, nil,
				},
				&TreeNode{
					4, nil, nil,
				},
			},
		},
		&TreeNode{
			1, &TreeNode{
				0, nil, nil,
			}, &TreeNode{
				8, nil, nil,
			},
		},
	}
	p := root.Left
	q := root.Right
	ans := lowestCommonAncestor(root, p, q)
	println(ans)

	fmt.Println(ans.Val)
}
