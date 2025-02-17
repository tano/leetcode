package dynamic

func climbStairs(n int) int {
	waysStorage := map[int]int{}
	return calculate(n, waysStorage)
}

func calculate(n int, storage map[int]int) int {
	if n == 1 {
		storage[1] = 1
		return 1
	}
	if n == 2 {
		storage[2] = 2
		return 2
	}
	if count, ok := storage[n]; !ok {
		count = calculate(n-1, storage) + calculate(n-2, storage)
		storage[n] = count
		return count
	} else {
		return count
	}
}
