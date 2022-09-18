package solved

func isIsomorphic(s string, t string) bool {
	charsAnalog := make(map[uint8]uint8)
	uniqueChars := make(map[uint8]uint8)
	if len(s) != len(t) {
		return false
	} else {
		for i := range s {
			x := s[i]
			y := t[i]
			if _, ok := charsAnalog[x]; ok {
				if charsAnalog[x] != y {
					return false
				}
			} else {
				if _, ok := uniqueChars[y]; ok {
					return false
				} else {
					charsAnalog[x] = y
					uniqueChars[y] = 1
				}
			}
		}
	}
	return true
}
