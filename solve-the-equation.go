package main

import (
	"fmt"
	"strings"
)

func solveEquation(equation string) string {
	expressions := strings.Split(equation, "=")

	x, c2 := eval(expressions[0])
	x2, c := eval(expressions[1])
	x -= x2
	c -= c2
	if x == 0 {
		if c == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}
	return fmt.Sprintf("x=%d", c/x)
}

func eval(expression string) (int, int) {
	x := 0
	c := 0
	isVar := false
	sign := 1
	i := 0
	t := 0
	xFlag := false
	for i < len(expression) {
		if expression[i] == '+' || expression[i] == '-' {
			if expression[i] == '+' {
				sign = 1
			} else {
				sign = -1
			}

			i++
		} else {
			t = 0
			for i < len(expression) && expression[i] != '+' && expression[i] != '-' {
				if expression[i] == 'x' {
					isVar = true
				} else {
					t = t*10 + int(expression[i]-'0')
					xFlag = true
				}
				i++
			}
			if t == 0 && !xFlag {
				t = 1
			}
			if isVar {
				x = x + sign*t
				isVar = false
			} else {
				c = c + sign*t
			}
			xFlag = false
		}
	}
	return x, c
}

func main() {
	//equation := "x+5-3+x=6+x-2"
	equation := "0x=0"
	ans := solveEquation(equation)
	fmt.Println(ans)
}
