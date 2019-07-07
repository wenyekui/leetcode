package main

import (
	"fmt"
	"math"
)

func findDuplicates2(nums []int) []int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if i != j {
			if nums[i] == nums[j] {
				i++
			} else {
				nums[i], nums[j] = nums[j], nums[i]
			}
		} else {
			i++
		}
	}

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		j := nums[i] - 1
		if i != j {
			result = append(result, nums[i])
		}
	}

	return result

}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		num = int(math.Abs(float64(num)))
		if nums[num-1] < 0 {
			res = append(res, num)
		} else {
			nums[num-1] *= -1
		}
	}
	return res
}

// fmt.Println(nums)
func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
	nums = []int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6}
	fmt.Println(findDuplicates(nums))
	println("***********************")
	nums = []int{3, 11, 8, 16, 4, 15, 4, 17, 14, 14, 6, 6, 2, 8, 3, 12, 15, 20, 20, 5}

	fmt.Println(findDuplicates(nums))

	println("***********************")
	nums = []int{27, 40, 6, 21, 14, 36, 10, 19, 44, 10, 41, 26, 39, 20, 25, 19, 14, 7, 29, 27, 40, 38, 11, 44, 4, 6, 48, 39, 9, 13, 7, 45, 41, 23, 31, 8, 24, 1, 3, 5, 28, 11, 49, 29, 18, 4, 38, 32, 24, 15}
	fmt.Println(findDuplicates(nums))
}
