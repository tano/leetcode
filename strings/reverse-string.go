package strings

func reverseString(s []byte) {
	increasing := 0
	decreasing := len(s) - 1
	for increasing <= decreasing {
		s[increasing], s[decreasing] = s[decreasing], s[increasing]
		increasing++
		decreasing--
	}
}
