package strings

import (
	"slices"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var candidates []string
	for i := 0; i < len(strs); i++ {
		candidate := strs[i]
		for j := len(candidate) - 1; j >= 0; j-- {
			prefixFor := isPrefixFor(candidate, strs)
			if prefixFor {
				candidates = append(candidates, candidate)
				break
			}
			candidate = candidate[:j]
		}
	}
	slices.Sort(candidates)
	if len(candidates) < 1 {
		return ""
	}
	return candidates[0]
}

func isPrefixFor(prefix string, strs []string) bool {
	for i := 0; i < len(strs); i++ {
		if strings.Index(strs[i], prefix) != 0 {
			return false
		}
	}
	return true
}
