package array

func plusOne(digits []int) []int {
	var result []int
	added := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			added = true
			break
		} else {
			digits[i] = 0
		}
	}
	if added {
		result = digits
	} else {
		result = append([]int{1}, digits...)
	}
	return result
}
