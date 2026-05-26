package main

import (
	"fmt"
	"math"
)

func main() {
	stocks := []int{10, 2, 1, 5, 7, 1}
	fmt.Println(BuySellStocks(stocks))
}

func BuySellStocks(stocks []int) int {
	res := 0
	l, r := 0, 1

	for r < len(stocks) {
		rs, ls := stocks[r], stocks[l]
		if ls < rs {
			res = int(math.Max(float64(res), float64(rs-ls)))
		} else {
			l = r
		}
		r++
	}
	return res
}
