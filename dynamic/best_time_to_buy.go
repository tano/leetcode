package dynamic

import "math"

func maxProfit(prices []int) int {
	sellPrice := 0
	buyPrice := math.MaxInt
	buyPtr, sellPtr := 0, len(prices)-1
	profit := 0

	for buyPtr <= sellPtr {
		if prices[buyPtr] < buyPrice {
			buyPrice = prices[buyPtr]
		}
		if prices[sellPtr] > sellPrice {
			sellPrice = prices[sellPtr]
		}
		potentialProfit := sellPrice - buyPrice
		if potentialProfit > profit {
			profit = potentialProfit
		}
		buyPtr++
		sellPtr--
	}

	return profit
}
