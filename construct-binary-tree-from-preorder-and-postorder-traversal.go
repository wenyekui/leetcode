package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	length := len(pre)
	if length == 0 {
		return nil
	}

	root := &TreeNode{pre[0], nil, nil}
	pre = pre[1:length]
	post = post[0 : length-1]

	length = len(pre)
	if length > 0 {
		pos := 0
		for i, v := range post {
			if pre[0] == v {
				pos = i
				break
			}
		}
		left := constructFromPrePost(pre[0:pos+1], post[0:pos+1])
		right := constructFromPrePost(pre[pos+1:length], post[pos+1:length])
		root.Left = left
		root.Right = right
	}
	return root
}

func main() {
	fmt.Println("")
}
