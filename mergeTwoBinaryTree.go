package LeetCode

func mergeTrees(root1 *BinaryTreeNode, root2 *BinaryTreeNode) *BinaryTreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	root := &BinaryTreeNode{0, nil, nil}
	if root1 != nil && root2 != nil {
		root.Val = root2.Val + root1.Val
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
	} else if root1 == nil {
		root.Val += root2.Val
		root.Left = mergeTrees(nil, root2.Left)
		root.Right = mergeTrees(nil, root2.Right)
	} else {
		root.Val += root1.Val
		root.Left = mergeTrees(nil, root1.Left)
		root.Right = mergeTrees(nil, root1.Right)
	}
	return root
}
