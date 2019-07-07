package main

import (
	"fmt"
)

func toHex(num int) string {
	unum := uint32(num)
	if unum == 0 {
		return "0"
	}
	res := ""
	for unum != 0 {
		q := unum % 16
		if q < 10 {
			res = fmt.Sprintf("%d", q) + res
		} else {
			res = fmt.Sprintf("%c", q-10+'a') + res
		}
		unum = unum / 16
	}
	return res
}

func main() {
	fmt.Println(toHex(-1))
	fmt.Println(toHex(26))
}
