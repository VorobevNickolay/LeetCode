package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	head_copy := head
	for head_copy != nil {
		if head_copy.Next == nil {
			break
		}
		head = head.Next
		head_copy = head_copy.Next.Next
	}
	return head
}
func main() {

}
