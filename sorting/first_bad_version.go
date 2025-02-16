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
	var versions []int
	for i := range n {
		versions = append(versions, i+1)
	}
	return findInSlice(versions)
}

func findInSlice(versions []int) int {
	if len(versions) == 1 {
		return versions[0]
	}
	if len(versions) == 2 {
		if isBadVersion(versions[0]) {
			return versions[0]
		} else {
			return versions[1]
		}
	}
	middle := len(versions) / 2
	if isBadVersion(versions[middle]) {
		if !isBadVersion(versions[middle-1]) {
			return versions[middle]
		} else {
			return findInSlice(versions[:middle])
		}
	} else {
		if isBadVersion(versions[middle+1]) {
			return versions[middle+1]
		} else {
			return findInSlice(versions[middle:])
		}
	}
}
