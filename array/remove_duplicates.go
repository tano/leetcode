package main

import "fmt"

func removeDuplicates(nums []int) int {
	current := 0
	curLen := len(nums)

	for {
		fmt.Printf("current: %v, array : %v, cur len: %v\n", current, nums, curLen)
		if current == curLen-1 {
			break
		}
		if nums[current] == nums[current+1] {
			leftShift(nums, current)
			curLen--
		} else {
			current++
		}

	}

	return current + 1
}

func leftShift(nums []int, pos int) {
	for i := pos; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

func main() {
	// [1,1,2]
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r := removeDuplicates(a)
	fmt.Printf("result: %v, array : %v", r, a[:r])
}
