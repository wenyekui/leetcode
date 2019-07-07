package main

import (
	"fmt"
	"math/rand"
	//	"time"
)

type RandomizedSet struct {
	mp   map[int]int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	//	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		map[int]int{},
		[]int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mp[val]; ok {
		return false
	}

	this.mp[val] = len(this.nums) - 1
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.mp[val]; !ok {
		return false
	}

	pos := this.mp[val]
	lastVal := this.nums[len(this.nums)-1]
	this.nums[pos] = lastVal
	this.mp[lastVal] = pos

	delete(this.mp, val)
	this.nums = this.nums[:len(this.nums)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {

	obj := Constructor()
	param_1 := obj.Insert(1)
	param_2 := obj.Remove(2)
	param_1 = obj.Insert(2)
	param_3 := obj.GetRandom()
	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
}
