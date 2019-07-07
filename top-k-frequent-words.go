package main

import (
	"container/heap"
	"fmt"
)

type wordHeap struct {
	words   []string
	counter map[string]int
}

func (h *wordHeap) Len() int {
	return len(h.words)
}

func (h *wordHeap) Less(i, j int) bool {
	if h.counter[h.words[i]] == h.counter[h.words[j]] {
		return h.words[i] < h.words[j]
	}
	return h.counter[h.words[i]] > h.counter[h.words[j]]
}

func (h *wordHeap) Swap(i, j int) {
	h.words[i], h.words[j] = h.words[j], h.words[i]
}

func (h *wordHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.words = append(h.words, x.(string))
}

func (h *wordHeap) Pop() interface{} {
	old := h.words
	n := len(old)
	x := old[n-1]
	h.words = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	counter := map[string]int{}
	for _, word := range words {
		counter[word] += 1
	}
	words = []string{}
	for word, _ := range counter {
		words = append(words, word)
	}

	h := &wordHeap{words, counter}
	heap.Init(h)
	ans := []string{}
	for k != 0 {
		ans = append(ans, (heap.Pop(h).(string)))
		k--
	}
	return ans
}

func main() {
	words := []string{
		"i", "love", "leetcode", "i", "love", "coding",
	}

	k := 2
	words = []string{
		"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is",
	}
	k = 4
	ans := topKFrequent(words, k)
	fmt.Printf("%v\n", ans)

func topKFrequent(words []string, k int) []string {
    counter := map[string]int{}
    for _, word := range words {
        counter[word] += 1
    }
    words = []string{}
    for word, _ := range counter {
        words = append(words, word)
    }

    h := &wordHeap{words, counter}
    heap.Init(h)
    ans := []string{}
    for k != 0 {
        ans = append(ans, (heap.Pop(h).(string)))
        k--
    }
    return ans
}}
