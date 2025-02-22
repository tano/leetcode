package dynamic

import "math"

type key struct {
	i, j int
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum := math.MinInt
	cache := map[key]int{}
	for i := 0; i < len(nums); i++ {
		for j := len(nums); j >= i; j-- {
			currentSum := calculateSum(i, j, nums, cache)
			if currentSum > sum {
				sum = currentSum
			}
		}
	}
	return sum
}

func calculateSum(begin int, end int, nums []int, cache map[key]int) int {
	k := key{i: begin, j: end}
	sum, ok := cache[k]
	if ok {
		return sum
	} else {
		result := 0
		for i := begin; i < end; i++ {
			result += nums[i]
			cache[key{begin, i}] = result
		}
		cache[k] = result
		return result
	}
}
