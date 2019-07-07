package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := nums[0]
	maxI := 0
	for i, num := range nums {
		if num > max {
			maxI = i
			max = num
		}
	}

	left := constructMaximumBinaryTree(nums[0:maxI])
	right := constructMaximumBinaryTree(nums[maxI+1 : len(nums)])
	root := &TreeNode{
		Val:   max,
		Left:  left,
		Right: right,
	}
	return root
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	fmt.Println(root.Val)
	fmt.Println(root.Left.Val)
	fmt.Println(root.Right.Val)
}
