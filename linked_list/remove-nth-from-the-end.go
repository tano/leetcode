package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	num := 0
	for cur != nil {
		num++
		cur = cur.Next
	}
	target := num - n
	cur = head
	for i := 0; i < target-1; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	} else {
		if num-1 == target {
			head = head.Next
		}
	}
	return head
}
