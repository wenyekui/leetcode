package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	if N == 1 {
		return []*TreeNode{&TreeNode{0, nil, nil}}
	}
	res := make([]*TreeNode, 0)

	for i := 1; i < N; i += 2 {
		leftNodes := allPossibleFBT(i)
		rightNodes := allPossibleFBT(N - i - 1)
		for _, left := range leftNodes {
			for _, right := range rightNodes {
				res = append(res, &TreeNode{0, left, right})
			}
		}

	}
	return res
}

func main() {

	fmt.Println(len(allPossibleFBT(7)))
}
