package solved

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	} else {
		i, j := 0, 0
		for i < len(s) && j < len(t) {
			if s[i] == t[j] {
				i++
			}
			j++
		}
		return i == len(s)
	}
}
