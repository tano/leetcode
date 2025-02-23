package dynamic

import "math"

func maxSubArray(nums []int) int {
	bestSum := math.MinInt32
	currentSum := 0
	for _, num := range nums {
		currentSum = max(num, currentSum+num)
		bestSum = max(bestSum, currentSum)
	}
	return bestSum
}
