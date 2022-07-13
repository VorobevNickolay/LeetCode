package LeetCode

func detectCycle(head *ListNode) int {
	if head != nil {
		m := make(map[*ListNode]bool)
		for head.Next != nil {
			if _, ok := m[head]; ok {
				return head.Val
			}
			m[head] = true
			head = head.Next
		}
	}
	return -1
}
