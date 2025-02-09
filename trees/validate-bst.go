package trees

import "math"

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val < min || node.Val > max {
		return false
	}
	return isValid(node.Left, min, node.Val-1) && isValid(node.Right, node.Val+1, max)
}
