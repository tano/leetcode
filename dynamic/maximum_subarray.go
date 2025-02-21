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
		for j := len(nums); j > i; j-- {
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
		candidate := nums[begin:end]
		result := 0
		for i := 0; i < len(candidate); i++ {
			result += candidate[i]
		}
		cache[k] = result
		return result
	}
}
