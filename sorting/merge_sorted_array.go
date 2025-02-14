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
	for idx1 < m {
		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else {
			shiftRight(nums1, idx1)
			nums1[idx1] = nums2[idx2]
			idx2++
			idx1++
		}
		if idx1 == m {
			idx1++
		}
	}
	for idx2 < n {
		nums1[idx1] = nums2[idx2]
		idx1++
		idx2++
	}
}

func shiftRight(nums []int, idx int) {
	for i := len(nums) - 1; i > idx; i-- {
		nums[i] = nums[i-1]
	}
}
