package dynamic

import "math"

func maxProfit(prices []int) int {
	buyingDay, sellingDay := 0, 0
	profit := 0
	maxSellPrice, minBuyPrice := 0, math.MaxInt
	stockBought := false

	for i := 0; i < len(prices); i++ {
		price := prices[i]

		if price < minBuyPrice {
			minBuyPrice = price
			if !stockBought {
				buyingDay = i
			}
		} else {
			if !stockBought {
				stockBought = true
				maxSellPrice = prices[buyingDay]
				sellingDay = buyingDay
			}
		}

		if price > maxSellPrice {
			maxSellPrice = price
			sellingDay = i
		} else {
			if stockBought {
				stockBought = false
				profit = prices[sellingDay] - prices[buyingDay]
			}
		}

		if i == len(prices)-1 && stockBought {
			return price - prices[buyingDay]
		}
	}

	if profit > 0 {
		return profit
	}

	return 0
}
