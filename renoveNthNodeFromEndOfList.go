package LeetCode

func removeNthFromEnd(head *ListNode, n int) []int {
	c := head
	c1 := head
	for i := 0; i < n; i++ {
		c = c.Next
	}
	if c == nil {
		return ListToArray(c1.Next)
	}
	for c.Next != nil {
		head = head.Next
		c = c.Next
	}
	if n == 1 {
		head.Next = nil
	} else {
		head.Next = head.Next.Next
	}
	return ListToArray(c1)
}
