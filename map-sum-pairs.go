package main

import (
	"fmt"
)

type MapSum struct {
	Nodes map[byte]*MapSum
	Value int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		Nodes: make(map[byte]*MapSum),
		Value: 0,
	}
}

func (this *MapSum) Insert(key string, val int) {
	if len(key) == 0 {
		this.Value = val
		return
	}

	node := this.Nodes[key[0]]
	if node == nil {
		tmp := Constructor()
		node = &tmp
		this.Nodes[key[0]] = node
	}

	node.Insert(key[1:len(key)], val)
}

func (this *MapSum) Sum(prefix string) int {

	if len(prefix) == 0 {
		sum := this.Value
		for _, node := range this.Nodes {
			sum += node.Sum(prefix)
		}
		return sum
	} else {
		node := this.Nodes[prefix[0]]
		if node != nil {
			return node.Sum(prefix[1:len(prefix)])
		}
	}
	return 0
}

func main() {

	obj := Constructor()
	obj.Insert("applie", 3)
	// obj.Insert("app", 2)
	param_2 := obj.Sum("ap")

	fmt.Println(param_2)
}
