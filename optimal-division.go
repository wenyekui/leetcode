package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func optimalDivision(nums []int) string {
	n := len(nums)
	MAX := [10][10]float32{}
	MIN := [10][10]float32{}
	MAX_POS := [10][10]int{}
	MIN_POS := [10][10]int{}
	for i := 0; i < n; i++ {
		MIN[i][i] = float32(nums[i])
		MAX[i][i] = float32(nums[i])
	}

	for d := 1; d < n; d++ {
		for i := 0; i+d < n; i++ {
			max := MAX[i][i] / MIN[i+1][i+d]
			min := MIN[i][i] / MAX[i+1][i+d]
			p1 := i
			p2 := i
			for j := i + 1; j < i+d; j++ {
				t1 := MAX[i][j] / MIN[j+1][i+d]
				t2 := MIN[i][j] / MAX[j+1][i+d]
				if t1 > max {
					max = t1
					p1 = j
				}
				if t2 < min {
					min = t2
					p2 = j
				}

			}
			MAX[i][i+d] = max
			MIN[i][i+d] = min
			MAX_POS[i][i+d] = p1
			MIN_POS[i][i+d] = p2
		}
	}
	left := []int{}
	right := []int{}
	var f func(int, int, bool)
	f = func(start int, end int, ismax bool) {
		if start >= end {
			return
		}
		var k int
		if ismax {
			k = MAX_POS[start][end]
		} else {
			k = MIN_POS[start][end]
		}
		if k != end-1 {
			left = append(left, k+1)
			right = append(right, end)
		}
		f(start, k, ismax)
		f(k+1, end, !ismax)
	}
	f(0, n-1, true)

	sort.Ints(left)
	sort.Ints(right)
	var bf bytes.Buffer
	left_i := 0
	right_i := 0
	for i, num := range nums {
		if left_i < len(left) && i == left[left_i] {
			bf.WriteString("(")
			left_i++
		}
		bf.WriteString(strconv.Itoa(num))
		if right_i < len(right) && i == right[right_i] {
			bf.WriteString(")")
			right_i++
		}

		if i != n-1 {
			bf.WriteString("/")
		}
	}
	return bf.String()
}

func main() {
	nums := []int{1000, 100, 10, 2}
	// nums = []int{1000, 1, 100000, 1}
	nums = []int{6, 2, 3, 4, 5}
	fmt.Println(optimalDivision(nums))
}

// println("''''''")
// println(int(MAX[0][n-1]))
// println(MIN[1][n-1])
// println(int(6 / MIN[1][n-1]))
// println(MAX_POS[0][n-1])
// println(MIN_POS[1][n-1])
// println(MIN_POS[1][n-2])
// println("''''''")
// fmt.Println(left)
// fmt.Println(right)
