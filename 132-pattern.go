package main

import (
	"fmt"
	"sort"
)

func find132pattern(nums []int) bool {
	P := make([]int, len(nums))
	for i := 0; i < len(P); i++ {
		P[i] = i
	}

	sort.Slice(P, func(i, j int) bool {
		return nums[P[i]] < nums[P[j]]
	})

	i, j := 0, len(P)-1
	for j-i >= 2 {
		println(i, P[i], nums[P[i]])
		println(j, P[j], nums[P[j]])
		println("--------------")
		if P[j] < P[i] {
			j--
			continue
		}
		if P[j] == j {
			j--
			continue
		}
		if P[j] == i+1 && P[j-1] == i {
			j = j - 2
			continue
		}
		if P[i] > P[j] {
			i++
			continue
		}
		if nums[P[j]] == nums[P[i]] {
			return false
		}

		println(i, P[i], nums[P[i]])
		println(j, P[j], nums[P[j]])

		for k := i + 1; k < j; k++ {

			println(k, P[k], nums[P[k]])
			if P[k] > P[j] && nums[P[k]] != nums[P[i]] && nums[P[k]] != nums[P[j]] {
				return true
			}
		}
		return false
	}
	return false
}

func main() {
	nums := []int{
		1, 2, 3, 4,
	}

	nums = []int{
		3, 5, 0, 3, 4,
	}
	x := find132pattern(nums)
	fmt.Println(x)
}
