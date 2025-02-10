package trees

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareNodes(root.Left, root.Right)
}

func compareNodes(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		return compareNodes(a.Left, b.Right) && compareNodes(a.Right, b.Left)
	}
	if a == nil || b == nil {
		return false
	}
	return true
}
