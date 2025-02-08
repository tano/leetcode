package trees

func maxDepth(root *TreeNode) int {
	var result int
	result = deeper(root, result)
	return result
}

func deeper(node *TreeNode, num int) int {
	if node == nil {
		return num
	}
	num += 1
	left := deeper(node.Left, num)
	right := deeper(node.Right, num)
	if left > right {
		return left
	} else {
		return right
	}
}
