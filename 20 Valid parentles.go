package main

import "testing"

func main() {

}

func isValid(s string) bool {
	mOE := map[rune]int{
		'(': 0,
		')': 0,
		'{': 0,
		'}': 0,
		'[': 0,
		']': 0,
	}
	for _, char := range s {
		if _, ok := mOE[char]; ok {
			mOE[char]++
		}
	}
	if mOE['('] == mOE[')'] && mOE['['] == mOE[']'] && mOE['{'] == mOE['}'] {
		return true
	}
	return false
}
func TestIsValid(t *testing.T) {
	got := isValid("[]{}")
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
