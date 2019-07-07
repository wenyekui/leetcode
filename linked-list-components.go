package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	M := make(map[int]bool)
	for _, v := range G {
		M[v] = true
	}
	ans := 1
	start := false
	for node := head; node != nil && len(M) != 0; node = node.Next {
		if M[node.Val] {
			if !start {
				start = true
			}
			delete(M, node.Val)
		} else {
			if start {
				ans++
				start = false
			}
		}
	}
	return ans
}

func main() {
	fmt.Println("")
}
