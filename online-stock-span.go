package main

import (
	"fmt"
)

type StockSpanner struct {
	Prices []int
}

func Constructor() StockSpanner {
	return StockSpanner{Prices: []int{}}
}

func (this *StockSpanner) Next(price int) int {
	ans := 0
	this.Prices = append(this.Prices, price)
	for i := len(this.Prices) - 1; i >= 0; i++ {
		if this.Prices[i] <= price {
			ans += 1
		} else {
			break
		}
	}
	return ans
}

func main() {
	stockSpanner := Constructor()
	println(stockSpanner.Next(100))
	fmt.Println("")
}
