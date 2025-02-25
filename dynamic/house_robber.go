package dynamic

func rob(nums []int) int {
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	sumEven, sumOdd := 0, 0
	for i := 0; i < len(nums); i += 1 {
		if i%2 == 0 {
			sumEven += nums[i]
		} else {
			sumOdd += nums[i]
		}
	}
	return max(sumEven, sumOdd)
}
