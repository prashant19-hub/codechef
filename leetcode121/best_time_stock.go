package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		profit := price - minPrice

		if profit > maxProfit {
			maxProfit = profit
		}

		if price < minPrice {
			minPrice = price
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	result := maxProfit(prices)

	fmt.Println("Maximum profit:", result)
}
