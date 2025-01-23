package array

func moveZeroes(nums []int) {
	var tmp []int
	for _, n := range nums {
		if n != 0 {
			tmp = append(tmp, n)
		}
	}

	for i := 0; i < len(nums); i++ {
		if i > len(tmp)-1 {
			nums[i] = 0
		} else {
			nums[i] = tmp[i]
		}
	}

}
