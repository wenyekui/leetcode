package main

import (
	"fmt"
)

func countBattleships(board [][]byte) int {
	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}

	n := len(board)
	m := len(board[0])
	count := 0
	for i := 0; i < n; i++ {
		marked := false
		for j := 0; j < m; j++ {
			if board[i][j] == '.' && marked {
				count++
				marked = false
				continue
			}
			if board[i][j] == 'X' {
				if !marked && (i == 0 || (i != 0 && board[i-1][j] != 'X')) {
					marked = true
				}
			}
		}
		if marked {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("")
}
