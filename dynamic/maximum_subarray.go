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
	cache := make(map[key]int, pow(len(nums), 2))
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
			curKey := key{begin, i}
			v, curOK := cache[curKey]
			if curOK {
				return v
			}
			result += nums[i]
			cache[curKey] = result
		}
		cache[k] = result
		return result
	}
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
