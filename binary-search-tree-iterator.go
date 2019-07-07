package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	vals []int
	cur  int
}

func getVals(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := getVals(root.Left)
	ans = append(ans, root.Val)
	return append(ans, (getVals(root.Right))...)

}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{getVals(root), -1}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.cur += 1
	return this.vals[this.cur]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.cur+1 < len(this.vals)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	fmt.Println("")
}
