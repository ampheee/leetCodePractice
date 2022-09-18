package solved

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindromeTest(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
	}
	var p *ListNode
	for s != nil {
		s.Next, p, s = p, s, s.Next
	}
	for p != nil {
		if head.Val != p.Val {
			return false
		}
		head, p = head.Next, p.Next
	}
	return true
}

func canConstructTest(ransomNote string, magazine string) bool {
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
