package linked_list

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func parseToListNodes(s string) *ListNode {
	numbers := strings.Split(s, ",")
	var head *ListNode
	var cur = &ListNode{}
	if len(numbers) == 1 {
		return head
	}
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
