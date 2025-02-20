package dynamic

import "math"

func maxProfit(prices []int) int {
	profit := 0
	minBuyPrice := math.MaxInt
	for buyPtr := 0; buyPtr < len(prices)-1; buyPtr++ {
		if prices[buyPtr] < minBuyPrice {
			minBuyPrice = prices[buyPtr]
		} else {
			continue
		}
		for sellPtr := buyPtr + 1; sellPtr < len(prices); sellPtr++ {
			if prices[sellPtr] > prices[buyPtr] {
				potentialProfit := prices[sellPtr] - prices[buyPtr]
				if potentialProfit > profit {
					profit = potentialProfit
				}
			}
		}
	}
	return profit
}
