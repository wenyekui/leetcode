package main

import (
	"fmt"
)

func canTransform(start string, end string) bool {
	N := len(start)

	stack := make([]int, 0, N)
	for i := 0; i < N; i++ {
		if start[i] == 'X' && end[i] == 'L' {

			stack = append(stack, 0)

		} else if start[i] == 'R' && end[i] == 'X' {

			stack = append(stack, 1)

		} else if start[i] == 'L' && end[i] == 'X' {
			if len(stack) != 0 && stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if start[i] == 'X' && end[i] == 'R' {
			if len(stack) != 0 && stack[len(stack)-1] == 1 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		} else if start[i] == end[i] {
			if start[i] == 'X' {
				continue
			}
			if start[i] == 'L' && len(stack) != 0 && stack[len(stack)-1] == 0 {
				continue
			}
			if start[i] == 'R' && len(stack) != 0 && stack[len(stack)-1] == 1 {
				continue
			}

			if len(stack) != 0 {
				return false
			}
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("")
}

//
//  class Solution(object):
//      def canTransform(self, start, end):
//          """
//          :type start: str
//          :type end: str
//          :rtype: bool
//          """
//          i, j = 0, 0
//          N = len(start)
//          while i < N and j < N:
//              while i < N - 1 and start[i] == 'X':
//                  i += 1
//              while j < N - 1 and end[j] == 'X':
//                  j += 1
//              if start[i] != end[j]:
//                  return False
//              if start[i] == 'L' and i < j:
//                  return False
//              if start[i] == 'R' and i > j:
//                  return False
//              i += 1
//              j += 1
//          return True
//
