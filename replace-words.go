package main

import (
	"fmt"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	result := make([]string, 0)
	words := strings.Split(sentence, " ")
	for _, word := range words {
		for _, root := range dict {
			if strings.HasPefix(word, root) {
				result = append(result, root)
			} else {
				result = append(result, word)
			}
		}
	}
	return strings.Join(result, " ")
}

func main() {

	fmt.Println("")
}
