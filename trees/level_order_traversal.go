package trees

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	calculated := calculateLevelOrder(root, result, 0)
	return calculated
}

func calculateLevelOrder(root *TreeNode, result [][]int, currentLevel int) [][]int {
	if len(result) == currentLevel {
		result = append(result, []int{})
	}
	result[currentLevel] = append(result[currentLevel], root.Val)
	if root.Left != nil {
		result = calculateLevelOrder(root.Left, result, currentLevel+1)
	}
	if root.Right != nil {
		result = calculateLevelOrder(root.Right, result, currentLevel+1)
	}
	return result
}
