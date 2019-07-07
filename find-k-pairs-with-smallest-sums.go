package main

import (
	"fmt"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}
	i, j, ii, jj := 0, 0, 0, 0
	ans := make([][]int, 0, k)
	for k != 0 {
		if i == ii && j == jj {
			ii = i + 1
			jj = j + 1

			ans = append(ans, []int{nums1[i], nums2[j]})
			k--
			continue
		}
		if i == len(nums1)-1 {
			for x := jj; x < len(nums2) && k > 0; x++ {
				ans = append(ans, []int{nums1[i], nums2[x]})
				k--
			}
			break
		}
		if j == len(nums2)-1 {
			for x := ii; x < len(nums1) && k > 0; x++ {
				ans = append(ans, []int{nums1[x], nums2[j]})
				k--
			}
			break
		}

		a, b := 0, 0
		if nums1[i]+nums2[jj] < nums1[ii]+nums2[j] {
			a = i
			b = jj
			jj += 1
		} else {
			a = ii
			b = j
			ii += 1
		}

		if ii >= len(nums1) {
			j += 1
			jj = j
			i += 1
			ii = i
		} else if jj >= len(nums2) {
			i += 1
			ii = i
			j += 1
			jj = j
		}
		ans = append(ans, []int{nums1[a], nums2[b]})
		k--

	}
	return ans
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	// ans := kSmallestPairs(nums1, nums2, k)
	// fmt.Printf("%v\n", ans)

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2

	nums1 = []int{1, 2, 3, 4}
	nums2 = []int{4}
	k = 10

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 10
	ans := kSmallestPairs(nums1, nums2, k)
	fmt.Printf("%v\n", ans)
}
