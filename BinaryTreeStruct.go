package LeetCode

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func binaryTreeHeight(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	l := binaryTreeHeight(root.Left)
	r := binaryTreeHeight(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
