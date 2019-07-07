package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	N := len(accounts)
	merged := make([]bool, N)
	ans := [][]string{}
	for i, account := range accounts {
		if merged[i] {
			continue
		}

		account[1:len(account)]

	}
}

func main() {
	fmt.Println("")
}
