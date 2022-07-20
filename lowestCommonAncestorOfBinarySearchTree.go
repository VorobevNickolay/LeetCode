package LeetCode

func lowestCommonAncestor(root, p, q *BinaryTreeNode) *BinaryTreeNode {
	for {
		if root.Val > p.Val && root.Val > q.Val && root.Left != nil {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val && root.Right != nil {
			root = root.Right
		} else {
			break
		}
	}
	return root

}
