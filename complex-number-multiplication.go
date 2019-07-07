package main

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	aParts := strings.Split(a, "+")
	bParts := strings.Split(b, "+")

	aR, _ := strconv.Atoi(aParts[0])
	aI, _ := strconv.Atoi(aParts[1][0 : len(aParts[1])-1])

	bR, _ := strconv.Atoi(bParts[0])
	bI, _ := strconv.Atoi(bParts[1][0 : len(bParts[1])-1])

	r := aR*bR - aI*bI
	i := aR*bI + bR*aI

	return fmt.Sprintf("%d+%di", r, i)
}

func main() {
	a := "1+1i"
	b := "1+1i"
	fmt.Println(complexNumberMultiply(a, b))
}
