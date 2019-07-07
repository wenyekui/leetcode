package main

import (
	"fmt"
)

type node struct {
	mid   int
	sum   int
	left  *node
	right *node
}

func build(nums []int) *node {
	fmt.Println(nums)
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	left := build(nums[0:mid])
	right := build(nums[mid+1:])

	sum := nums[mid]

	if left != nil {
		sum += left.sum
	}
	if right != nil {
		sum += right.sum
	}
	return &node{mid, sum, left, right}
}

type NumArray struct {
	root *node
}

func Constructor(nums []int) NumArray {
	return NumArray{build(nums)}
}

func (this *NumArray) Update(i int, val int) {
	dfs_update(this.root, i, val)
}

func dfs_update(root *node, i int, val int) {
	if root == nil {
		return
	}
	m := root.mid
	if i < m {
		dfs_update(root.left, i, val)
	} else if i > m {
		dfs_update(root.right, i-m, val)
	}
	root.sum = val
	if root.left != nil {
		root.sum += root.left.sum
	}
	if root.right != nil {
		root.sum += root.right.sum
	}
}

func dfs(root *node, i int, j int) int {
	m := root.mid
	if j < m {
		return dfs(root.left, i, j)
	}
	if i >= m {
		return dfs(root.right, i-m, j-m)
	}
	if root.mid == 0 {
		return root.sum
	}
	return dfs(root.left, i, m-1) + dfs(root.right, 0, j-m)
}

func (this *NumArray) SumRange(i int, j int) int {
	return dfs(this.root, i, j)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

func main() {
	fmt.Println("")

	nums := []int{1, 3, 5}

	obj := Constructor(nums)

	println(obj.root.mid)
	// s := obj.SumRange(0, 2)
	// println(s)

	// obj.Update(1, 2)
	// s = obj.SumRange(0, 2)

	// println(s)
}
