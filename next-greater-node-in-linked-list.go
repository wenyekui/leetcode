package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	stack := []int{}
	nums := []int{}
	for node := head; node != nil; node = node.Next {
		nums = append(nums, node.Val)
	}

	ans := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return ans
}

func main() {
	fmt.Println("")
}
