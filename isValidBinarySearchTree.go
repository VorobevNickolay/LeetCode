package LeetCode

import "math"

func isVBST(root *BinaryTreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val > max || root.Val < min {
		return false
	}
	a := isVBST(root.Left, min, root.Val-1)
	b := isVBST(root.Right, root.Val+1, max)
	return a && b
}
func isValidBST(root *BinaryTreeNode) bool {
	return isVBST(root, math.MinInt, math.MaxInt)
}
