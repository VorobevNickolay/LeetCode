package LeetCode

func reverseList(head *ListNode) []int {
	var prev, cur *ListNode
	if head != nil {
		cur = head
		for head.Next != nil {
			cur = head.Next
			head.Next = prev
			prev = head
			head = cur
		}
		head.Next = prev
	}
	return ListToArray(cur)
}
