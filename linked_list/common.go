package linked_list

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func parseToListNodes(s string) *ListNode {
	numbers := strings.Split(s, ",")
	var head *ListNode
	var cur = &ListNode{}
	for i := 0; i < len(numbers); i++ {
		sCurNum := numbers[i]
		curNum, err := strconv.Atoi(sCurNum)
		if err != nil {
			panic(err)
		}
		cur.Val = curNum
		if i == 0 {
			head = cur
		}
		if i < len(numbers)-1 {
			next := &ListNode{}
			cur.Next = next
			cur = next
		}
	}
	return head
}
