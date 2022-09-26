package easy

func longestCommonPrefix(strs []string) string {
	var min int
	for _, val := range strs {
		if len(val) > min {
			min = len(val)
		}
	}
	for i := 0; i < min-1; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] != strs[j][i] {
				return strs[j-1][0:i]
			}
		}
	}
	return ""
}
