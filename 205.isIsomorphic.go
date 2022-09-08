package main

import "reflect"

func isIsomorphic(s string, t string) bool {
	charsAnalog := make(map[uint8]uint8)
	uniqueChars := make(map[uint8]bool)
	if len(s) != len(t) {
		return false
	} else {
		for i := range s {
			x := s[i]
			y := t[i]
			if containsKey(charsAnalog, x) {
				if charsAnalog[x] != y {
					return false
				}
			} else {
				if containsKey(uniqueChars, y) {
					return false
				} else {
					charsAnalog[x] = y
					uniqueChars[y] = true
				}
			}
		}

	}
	return true
}

func containsKey(m, k interface{}) bool {
	v := reflect.ValueOf(m).MapIndex(reflect.ValueOf(k))
	return v != reflect.Value{}
}
