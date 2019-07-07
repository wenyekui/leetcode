package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	node := root
	for true {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		if k == 1 {
			return stack[len(stack)-1].Val
		}

		node = stack[len(stack)-1].Right
		stack = stack[0 : len(stack)-1]
		k--
	}
	return 0
}

func main() {

	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 4, Left: nil, Right: nil},
	}

	x := kthSmallest(root, 1)
	fmt.Println(x)
	root = &TreeNode{
		5,
		&TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{6, nil, nil},
	}
	x = kthSmallest(root, 3)
	fmt.Println(x)
}
