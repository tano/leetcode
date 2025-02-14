package sorting

func merge(nums1 []int, m int, nums2 []int, n int) {
	var idx1, idx2 int
	if m == 0 && n == 0 {
		return
	}
	if m == 1 && n == 0 {
		return
	}
	if m == 0 && n == 1 {
		nums1[idx1] = nums2[idx2]
		return
	}
	var result []int
	for idx1+idx2 < m+n {

		if idx1 < m && idx2 < n && nums1[idx1] < nums2[idx2] {
			result = append(result, nums1[idx1])
			idx1++
		} else if idx2 < n {
			result = append(result, nums2[idx2])
			idx2++
		} else if idx1 < m {
			result = append(result, nums1[idx1])
			idx1++
		}
	}
	for i, num := range result {
		nums1[i] = num
	}
}
