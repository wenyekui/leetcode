package main

import (
	"fmt"
)

func find(dsu map[byte]byte, x byte) byte {
	if _, ok := dsu[x]; !ok || dsu[x] == x {
		dsu[x] = x
	} else {
		dsu[x] = find(dsu, dsu[x])
	}
	return dsu[x]
}

func union(dsu map[byte]byte, x byte, y byte) {
	dsu[find(dsu, x)] = find(dsu, y)
}

func equationsPossible(equations []string) bool {

	dsu := make(map[byte]byte)
	for _, equation := range equations {
		if equation[1:3] == "==" {
			x := equation[0]
			y := equation[3]
			union(dsu, x, y)
		}
	}

	for _, equation := range equations {
		if equation[1:3] == "!=" {
			x := equation[0]
			y := equation[3]

			if find(dsu, x) == find(dsu, y) {
				return false
			}

		}
	}
	return true
}

func main() {
	equations := []string{"a==b", "b!=a"}
	ans := equationsPossible(equations)
	fmt.Println(ans)

	equations = []string{"a==b", "b!=c", "c==a"}
	ans = equationsPossible(equations)
	fmt.Println(ans)

	equations = []string{"a==b", "b==c", "c==a"}
	ans = equationsPossible(equations)
	fmt.Println(ans)

	equations = []string{"a==b", "e==c", "b==c", "a!=e"}
	ans = equationsPossible(equations)
	fmt.Println(ans)
}
