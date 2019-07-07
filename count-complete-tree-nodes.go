package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := 0
	for node := root.Left; node != nil; node = node.Left {
		lh++
	}
	rh := 0
	for node := root.Right; node != nil; node = node.Right {
		rh++
	}
	if lh == rh {
		return int(math.Pow(2, float64(lh))) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right)
}

func main() {
	fmt.Println("")
}
