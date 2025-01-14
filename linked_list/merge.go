package linked_list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curFirst, curSecond := list1, list2
	var curResult, curHead *ListNode

	for {
		var value int

		if curFirst == nil && curSecond == nil {
			break
		}

		if curFirst == nil {
			value = curSecond.Val
			curSecond = curSecond.Next
		} else if curSecond == nil {
			value = curFirst.Val
			curFirst = curFirst.Next
		} else {
			if curFirst.Val <= curSecond.Val {
				value = curFirst.Val
				curFirst = curFirst.Next
			} else {
				value = curSecond.Val
				curSecond = curSecond.Next
			}
		}

		if curResult == nil {
			curResult = &ListNode{Val: value}
			curHead = curResult
		} else {
			curResult.Val = value
		}
		if curFirst != nil || curSecond != nil {
			next := &ListNode{}
			curResult.Next = next
			curResult = curResult.Next
		}
	}

	return curHead
}
