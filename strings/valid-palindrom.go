package strings

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	filtered := strings.ToLower(filterAlphabet(s))
	if filtered == revString(filtered) {
		return true
	}
	return false
}

func filterAlphabet(s string) string {
	runes := []rune(s)
	var result []rune
	for _, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func revString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
