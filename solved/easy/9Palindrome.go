package easy

import "strconv"

func isPalindrome(x int) bool {
	flag := true
	stringInt := strconv.Itoa(x)
	for i := 0; i < len(stringInt)/2; i++ {
		if stringInt[i] != stringInt[len(stringInt)-i-1] {
			flag = false
			break
		}
	}
	return flag
}
