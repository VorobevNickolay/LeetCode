package LeetCode

func isPalindrome(head *ListNode) bool {
	head_copy := head
	head1 := head
	if head.Next == nil {
		return true
	}
	for head_copy != nil {
		if head_copy.Next == nil {
			break
		}
		head = head.Next
		head_copy = head_copy.Next.Next
	}
	if head == nil {
		return true
	}
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
	for cur != nil {
		if head1.Val != cur.Val {
			return false
		}
		head1 = head1.Next
		cur = cur.Next
	}
	return true
}
