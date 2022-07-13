package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToList(nums []int) *ListNode {
	m := make([]ListNode, len(nums))
	for i, num := range nums {
		m[i].Val = num
	}
	for i := 0; i < len(nums)-1; i++ {
		m[i].Next = &m[i+1]
	}
	return &m[0]
}

func ListToArray(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}

func ListWithCycle(nums []int, pos int) *ListNode {
	m := make([]ListNode, len(nums))
	for i, num := range nums {
		m[i].Val = num
	}
	for i := 0; i < len(nums)-1; i++ {
		m[i].Next = &m[i+1]
	}
	m[len(nums)-1].Next = &m[pos]
	return &m[0]
}
func ListTest(nums []int) bool {
	head := ArrayToList(nums)
	test := ListToArray(head)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != test[i] {
			return false
		}
	}
	return true
}
