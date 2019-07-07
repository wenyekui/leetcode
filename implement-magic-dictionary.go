package main

import (
	"fmt"
)

type MagicDictionary struct {
	dict map[string]map[rune]bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{make(map[string]map[rune]bool)}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, word := range dict {
		for i, r := range word {
			key := word[0:i] + "*" + word[i+1:len(word)]
			if this.dict[key] == nil {
				this.dict[key] = make(map[rune]bool)
			}
			this.dict[key][r] = true

		}
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	for i, r := range word {
		key := word[0:i] + "*" + word[i+1:len(word)]
		if this.dict[key] != nil {
			if len(this.dict[key]) > 1 || !this.dict[key][r] {
				return true
			}
		}
	}
	return false
}
func main() {
	fmt.Println("")
}
