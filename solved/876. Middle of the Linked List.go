package solved

func middleNode(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
	}
	return s
}
