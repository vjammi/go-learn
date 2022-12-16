package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 6, 4, 3, 1, 10}
	profit := maxProfit(prices)
	fmt.Print(profit)
}

func maxProfit(prices []int) int {

	var minPriceIndex int
	var priceOnMinPriceIndex int
	var profit int

	minPriceIndex = 0
	priceOnMinPriceIndex = math.MaxInt64
	profit = 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < priceOnMinPriceIndex {
			minPriceIndex = i
			priceOnMinPriceIndex = prices[i]
		}
		profit = int(math.Max(float64(profit), float64(prices[i]-prices[minPriceIndex])))
	}
	return profit
}
