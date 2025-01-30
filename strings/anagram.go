package strings

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if sortString(s) == sortString(t) {
		return true
	}
	return false
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
