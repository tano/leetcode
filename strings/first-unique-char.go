package strings

import "strings"

func firstUniqChar(s string) int {
	runes := []rune(s)
	checkStats := map[rune]int{}
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		if count, ok := checkStats[char]; ok {
			if count > 1 {
				continue
			}
			checkStats[char] += 1
		} else {
			checkStats[char] = 1
		}
		count := strings.Count(string(runes), string(char))
		if count == 1 {
			return i
		}
	}
	return -1
}
