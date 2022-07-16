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

/*
func ArrayToBinaryTree(arr []int) *BinaryTreeNode {
	var head, headCopy *BinaryTreeNode
	for i, num := range arr {
		if num != -1 {
			head.Val = num
		} else {
			head = nil
		}

	}
}
*/
