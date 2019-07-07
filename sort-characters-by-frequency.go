package main

import (
	"bytes"
	"fmt"
	"sort"
)

type Frequency struct {
	C rune
	N int
}

func frequencySort(s string) string {
	counter := map[rune]int{}

	for _, c := range s {
		counter[c]++
	}
	h := make([]*Frequency, 0, len(counter))
	for c, n := range counter {
		h = append(h, &Frequency{c, n})
	}

	sort.Slice(h, func(i, j int) bool {
		return h[i].N > h[j].N
	})
	var bf bytes.Buffer

	for _, x := range h {
		for i := 0; i < x.N; i++ {
			bf.WriteRune(x.C)
		}
	}
	return bf.String()
}

func main() {

	fmt.Println(frequencySort("tree"))
}
