package linked_list

import "math"

func hasCycle(head *ListNode) bool {
	markerValue := int(math.Pow(float64(10), float64(5))) + 1
	current := head
	for current != nil {
		if current.Val == markerValue {
			return true
		}
		current.Val = markerValue
		current = current.Next
	}
	return false
}
