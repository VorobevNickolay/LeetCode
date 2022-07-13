package LeetCode

func middleNode(head *ListNode) int {
	if head == nil {
		return -1
	}
	head_copy := head
	for head_copy != nil {
		if head_copy.Next == nil {
			break
		}
		head = head.Next
		head_copy = head_copy.Next.Next
	}
	return head.Val
}
