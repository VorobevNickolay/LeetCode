package LeetCode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pl := 0
	head := l1
	s := 0
	for true {
		s = l1.Val + l2.Val + pl
		l1.Val = s % 10
		pl = s / 10
		if l1.Next == nil {
			if l2.Next == nil {
				if pl != 0 {
					l1.Next = &ListNode{pl, nil}
				}
				return head
			}
			l1.Next = l2.Next
			l2.Next = nil
		}
		l1 = l1.Next
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
	}
	return head
}
