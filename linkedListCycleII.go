package LeetCode

func detectCycle(head *ListNode) *ListNode {
	if head != nil {
		m := make(map[*ListNode]bool)
		for head.Next != nil {
			if _, ok := m[head]; ok {
				return head
			}
			m[head] = true
			head = head.Next
		}
	}
	return nil
}
