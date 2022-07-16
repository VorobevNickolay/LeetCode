package LeetCode

func height(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	l := height(root.Left)
	r := height(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
func printLevel(root *BinaryTreeNode, level int) []int {
	var res []int
	if root == nil {
		return []int{}
	}
	if level == 0 {
		return []int{root.Val}
	} else if level > 0 {
		res = append(res, printLevel(root.Left, level-1)...)
		res = append(res, printLevel(root.Right, level-1)...)
	}
	return res
}
func levelOrder(root *BinaryTreeNode) [][]int {
	h := height(root)
	arr := make([][]int, h)
	for i := 0; i < h; i++ {
		arr[i] = printLevel(root, i)
	}
	return arr
}
