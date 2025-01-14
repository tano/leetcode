package linked_list

func isPalindrome(head *ListNode) bool {
	// corner cases
	if head == nil {
		// zero elements
		return false
	}
	if head.Next == nil {
		// single elements
		return true
	}
	if head.Next != nil && head.Next.Next == nil {
		// two elements
		if head.Val == head.Next.Val {
			return true
		} else {
			return false
		}
	}

	// base case
	var stack []int
	var cur *ListNode
	cur = head
	allVisited := false
	for !allVisited {
		stack = append(stack, cur.Val)
		if cur.Next == nil {
			allVisited = true
		}
		cur = cur.Next
	}
	i, j := 0, len(stack)-1
	for i < j {
		if stack[i] != stack[j] {
			return false
		}
		i++
		j--
	}
	return true
}
