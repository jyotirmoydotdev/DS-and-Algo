package main

import (
	"math"
)

// func main() {
// 	prices := []int{7, 6, 4, 3, 1}
// 	fmt.Println("MaxProfit:", maxProfit(prices))
// }

func maxProfit(prices []int) int {
	min := math.MaxUint32
	res := 0

	for _, price := range prices {
		if price > min {
			if price-min > res {
				res = price - min
			}
		} else {
			min = price
		}
	}
	return res
}
