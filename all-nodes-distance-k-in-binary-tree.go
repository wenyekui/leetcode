package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	ans := []int{}
	dfs(root, target, K, &ans)
	return ans
}

func dfs(root *TreeNode, target *TreeNode, k int, ans *[]int) int {
	if root == nil {
		return -1
	}

	if root == target {
		if k == 0 {
			*ans = append(*ans, root.Val)
		} else {
			collect(root.Left, k, 1, ans)
			collect(root.Right, k, 1, ans)
		}
		return 0
	}

	l := dfs(root.Left, target, k, ans)
	if l != -1 && l < k {
		if l == k-1 {
			*ans = append(*ans, root.Val)
		}
		collect(root.Right, k-l-1, 1, ans)
		return l + 1
	}

	l = dfs(root.Right, target, k, ans)
	if l != -1 && l < k {
		if l == k-1 {
			*ans = append(*ans, root.Val)
		}
		collect(root.Left, k-l-1, 1, ans)
		return l + 1
	}

	return -1
}

func collect(node *TreeNode, k int, d int, ans *[]int) {
	if node == nil {
		return
	}
	if k == d {
		*ans = append(*ans, node.Val)
		return
	}
	collect(node.Left, k, d+1, ans)
	collect(node.Right, k, d+1, ans)
}

func main() {
	fmt.Println("")
}
