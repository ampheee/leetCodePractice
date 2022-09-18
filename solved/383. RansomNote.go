package solved

func canConstruct(ransomNote string, magazine string) bool {
	m := map[rune]int{}
	for _, valRune := range magazine {
		m[rune(valRune)] += 1
	}
	for _, valRune := range ransomNote {
		if val, ok := m[valRune]; !ok || val == 0 {
			return false
		}
		m[valRune] -= 1
	}
	return true
}
