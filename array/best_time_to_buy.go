package main

import "math"

func maxProfit(prices []int) int {
	buyPrice := math.MaxInt32
	sellPrice := 0

	profit := 0
	bought := false

	if len(prices) < 2 {
		return 0
	}

	if len(prices) == 2 {
		if prices[1] > prices[0] {
			return prices[1] - prices[0]
		} else {
			return 0
		}
	}

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] <= buyPrice {
			buyPrice = prices[i]
			if prices[i+1] > buyPrice {
				bought = true
				if i == len(prices)-2 {
					sellPrice = prices[i+1]
					profit += sellPrice - buyPrice
				}
			}
		} else {
			sellPrice = prices[i]
			if (prices[i+1] < sellPrice) && bought {
				profit += sellPrice - buyPrice
				buyPrice = math.MaxInt32
			}
			if i == len(prices)-2 && buyPrice < sellPrice {
				sellPrice = prices[i+1]
				profit += sellPrice - buyPrice
			}
		}
	}

	return profit
}
