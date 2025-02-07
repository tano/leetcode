package linked_list

import "slices"

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var values []int
	cur := head
	for cur != nil {
		values = append(values, cur.Val)
		cur = cur.Next
	}
	slices.Reverse(values)
	var result = &ListNode{}
	cur = result
	for i, value := range values {
		cur.Val = value
		if i != len(values)-1 {
			next := &ListNode{}
			cur.Next = next
			cur = cur.Next
		}
	}
	return result
}
