package strings

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	trimmed := strings.TrimSpace(s)
	digitStarted := false
	isNegative := false
	isPositive := false
	result := ""
	if len(trimmed) == 0 {
		return 0
	}
	for _, char := range trimmed {
		if unicode.IsDigit(char) {
			digitStarted = true
			result += string(char)
		} else {
			if digitStarted {
				break
			}
			if char == '+' {
				if isNegative || isPositive {
					return 0
				}
				isPositive = true
				continue
			}
			if char == '-' {
				if isPositive || isNegative {
					return 0
				}
				isNegative = true
				continue
			}

			return 0
		}
	}
	if result == "" {
		return 0
	}
	atoi, err := strconv.Atoi(result)
	if err != nil {
		if isNegative {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	if isNegative {
		atoi = -atoi
		if atoi < math.MinInt32 {
			return math.MinInt32
		}
		return atoi
	}
	if atoi > math.MaxInt32 {
		return math.MaxInt32
	}
	return atoi
}
