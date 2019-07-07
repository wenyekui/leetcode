package main

import (
	"strconv"
)

type Parser struct {
	pos int
	s   string
}

func (this *Parser) readNumber() int {
	start := this.pos
	for this.s[this.pos] >= '0' && this.s[this.pos] <= '9' {
		this.pos++
	}
	res, _ := strconv.Atoi(this.s[start:this.pos])
	return res
}

func (this *Parser) readString() string {
	start := this.pos
	for this.pos < len(this.s) && (this.s[this.pos] >= 'a' && this.s[this.pos] <= 'z' || this.s[this.pos] >= 'A' && this.s[this.pos] <= 'Z') {
		this.pos++
	}
	return this.s[start:this.pos]
}

func (this *Parser) run() string {
	ans := ""
	for this.pos < len(this.s) && this.s[this.pos] != ']' {
		if this.s[this.pos] >= '0' && this.s[this.pos] <= '9' {
			num := this.readNumber()
			this.pos++
			r := this.run()
			for i := 0; i < num; i++ {
				ans += r
			}
			this.pos++
		} else {
			ans += this.readString()
		}
	}
	return ans
}

func decodeString(s string) string {
	parser := Parser{0, s}
	return parser.run()
}

func main() {
	s := "3[a]2[bc]"
	ans := decodeString(s)
	println(ans)

	s = "3[a2[c]]"
	ans = decodeString(s)
	println(ans)
}
