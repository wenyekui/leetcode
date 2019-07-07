package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	return construct(head, n)
}

func construct(head *ListNode, n int) *TreeNode {
	if n == 0 {
		return nil
	}

	if n == 1 {
		return &TreeNode{
			head.Val,
			nil,
			nil,
		}
	}
	count := 1
	node := head
	for node = head; count != (n+1)/2; node = node.Next {
		count++
	}

	tn := &TreeNode{
		node.Val,
		nil,
		nil,
	}
	tn.Left = construct(head, count-1)
	tn.Right = construct(node.Next, n-count)
	return tn
}

func main() {
	fmt.Println("")
}
