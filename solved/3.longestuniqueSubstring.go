package solved

import "strings"

func lengthOfLongestSubstring(s string) int {
	counter, finCount := 1, 0
	if len(s) == 1 {
		finCount = 1
	} else if len(s) == 0 {
		finCount = 0
	} else {
		tempString := ""
		tempString += string(s[0])
		for i := 1; i < len(s); i++ {
			if !strings.Contains(tempString, string(s[i])) {
				counter++
				tempString += string(s[i])
			} else {
				tempString = ""
				tempString += string(s[i])
				counter = 1
			}
			if counter > finCount {
				finCount = counter
			}
		}
	}
	return finCount
}
