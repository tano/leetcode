package array

func intersect(nums1 []int, nums2 []int) []int {
	nums1Stat := map[int]int{}
	nums2Stat := map[int]int{}
	var res []int
	for _, num := range nums1 {
		nums1Stat[num]++
	}
	for _, num := range nums2 {
		nums2Stat[num]++
	}
	for num, count := range nums1Stat {
		countInSecond := nums2Stat[num]
		repeatCount := 0
		if count > countInSecond {
			repeatCount = countInSecond
		} else {
			repeatCount = count
		}
		for range repeatCount {
			res = append(res, num)
		}
	}

	return res
}
