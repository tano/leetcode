package main

func rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}

	target := len(nums) - 1
	for range k - 1 {
		if target == 0 {
			target = len(nums) - 1
		} else {
			target--
		}
	}

	result := append(nums[target:], nums[:target]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}
}
