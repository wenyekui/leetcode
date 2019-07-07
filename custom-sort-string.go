package main

import (
	"fmt"
	"sort"
)

var M map[byte]int

type ByteSlice []byte

func (c ByteSlice) Len() int {
	return len(c)
}

func (c ByteSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ByteSlice) Less(i, j int) bool {
	return M[c[i]] < M[c[j]]
}

func customSortString(S string, T string) string {
	M = make(map[byte]int)
	for i, c := range []byte(S) {
		M[c] = i
	}
	bs := ByteSlice(T)
	sort.Sort(bs)
	return string(bs)
}

func main() {
	S := "cba"
	T := "abcd"

	fmt.Println(customSortString(S, T))
}
