package LeetCode

func preorder(root *TreeNode) []int {
	var a []int
	var b []int
	if root != nil {
		a = append(a, root.Val)
		for i := range root.Children {
			b = preorder(root.Children[i])
			a = append(a, b...)
		}
	}
	return a
}
