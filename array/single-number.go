package array

func singleNumber(nums []int) (result int) {
	var stats = make(map[int]int)
	for _, num := range nums {
		if _, ok := stats[num]; ok {
			stats[num]++
		} else {
			stats[num] = 1
		}
	}
	for key, value := range stats {
		if value == 1 {
			result = key
		}
	}
	return
}
