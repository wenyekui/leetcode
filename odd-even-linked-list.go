package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := head.Next
	is_odd := true

	for node := head.Next; node != nil; node = node.Next {
		if is_odd {
			even.Next = node.Next
			node.Next = even
			odd.Next = node
			odd = node
			node = even.Next
			even = even.Next
			is_odd = false
			continue
		} else {
			is_odd = true
		}
	}
	return head
}

func main() {
	fmt.Println("")
}
