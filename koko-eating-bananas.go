package main

import (
//	"fmt"
)

func minEatingSpeed(piles []int, H int) int {
	sum := 0
	for _, p := range piles {
		sum += p
	}
	avg := sum / H

	if sum%H != 0 {
		avg += 1
	}

	for true {
		xx := 0
		for _, p := range piles {
			x := p / avg
			if p%avg != 0 {
				x += 1
			}
			xx += x
		}
		if xx == H {
			break
		}
		avg += 1
	}

	return avg

}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	H := 6
	x := minEatingSpeed(piles, H)
	println(x)
}
