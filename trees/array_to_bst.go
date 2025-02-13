package trees

func sortedArrayToBST(nums []int) *TreeNode {
	var result = &TreeNode{}
	sliceToBST(nums, result)
	return result
}

func sliceToBST(nums []int, node *TreeNode) {
	numsLen := len(nums)
	if numsLen < 1 {
		return
	}
	if numsLen == 1 {
		node.Val = nums[0]
		return
	}
	if numsLen == 2 {
		node.Val = nums[1]
		node.Left = &TreeNode{Val: nums[0]}
		return
	}
	if numsLen == 3 {
		node.Left = &TreeNode{Val: nums[0]}
		node.Val = nums[1]
		node.Right = &TreeNode{Val: nums[2]}
		return
	}
	mid := numsLen / 2
	node.Val = nums[mid]
	node.Left = &TreeNode{}
	node.Right = &TreeNode{}
	sliceToBST(nums[:mid], node.Left)
	sliceToBST(nums[mid+1:], node.Right)
	return
}
