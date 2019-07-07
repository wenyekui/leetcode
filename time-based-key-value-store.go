package main

import (
	"fmt"
)

type Node struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]*Node
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		make(map[string][]*Node),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.m[key] == nil {
		this.m[key] = []*Node{}
	}
	this.m[key] = append(this.m[key], &Node{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if this.m[key] == nil || len(this.m[key]) == 0 || this.m[key][0].timestamp > timestamp {
		return ""
	}

	m := this.m[key]
	low := 0
	hight := len(m) - 1

	for low <= hight {
		mid := (low + hight) / 2
		if m[mid].timestamp == timestamp {
			return m[mid].value
		}
		if m[mid].timestamp > timestamp {
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}
	return m[low-1].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	fmt.Println("")
}
