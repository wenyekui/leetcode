package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i, x := range flowerbed {
		if x == 1 {
			continue
		}
		if (i == 0 || flowerbed[i-1] == 0) && (i+1 == len(flowerbed) || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
		if n == 0 {
			return true
		}
	}
	return false
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))

	flowerbed = []int{1, 0, 0, 0, 1}
	n = 2
	fmt.Println(canPlaceFlowers(flowerbed, n))

	flowerbed = []int{1, 0, 0, 0, 0, 1}
	n = 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}
