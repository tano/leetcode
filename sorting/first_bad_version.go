package sorting

type checker func(version int) bool

var defaultChecker checker = func(version int) bool { return version == 3 }

func isBadVersion(version int) bool {
	return defaultChecker(version)
}

func firstBadVersion(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		if isBadVersion(1) == false {
			return 2
		} else {
			return 1
		}
	}
	return find(1, n)
}

func find(startIndex, endIndex int) int {
	if startIndex == endIndex {
		return startIndex
	}
	if endIndex-startIndex == 1 {
		if isBadVersion(startIndex) {
			return startIndex
		} else {
			return endIndex
		}
	}
	middle := (endIndex + startIndex) / 2
	if isBadVersion(middle) {
		if !isBadVersion(middle) {
			return middle
		} else {
			return find(startIndex, middle)
		}
	} else {
		if isBadVersion(middle + 1) {
			return middle + 1
		} else {
			return find(middle+1, endIndex)
		}
	}
}
