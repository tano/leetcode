package strings

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	xAsStr := strconv.Itoa(x)
	bytes := []byte(xAsStr)
	reverseStringHelper(bytes)
	reversed := string(bytes)
	reversedInt, err := strconv.Atoi(reversed)
	if err != nil || reversedInt > math.MaxInt32 || reversedInt < math.MinInt32 {
		return 0
	}
	if negative {
		reversedInt = -reversedInt
	}
	return reversedInt
}

func reverseStringHelper(s []byte) {
	increasing := 0
	decreasing := len(s) - 1
	for increasing <= decreasing {
		s[increasing], s[decreasing] = s[decreasing], s[increasing]
		increasing++
		decreasing--
	}
}
