package linked_list

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "head = [3,2,0,-4], pos = 1",
			args: args{
				head: parsePotentiallyCycled("3,2,0,-4", 1),
			},
			want: true,
		},
		{
			name: "head = [1,2], pos = 0",
			args: args{
				head: parsePotentiallyCycled("1,2", 0),
			},
			want: true,
		},
		{
			name: "head = [1], pos = -1",
			args: args{
				head: parsePotentiallyCycled("1", -1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parsePotentiallyCycled(s string, pos int) *ListNode {
	head := parseToListNodes(s)
	if pos == -1 {
		return head
	}
	var curNode, nodeToPoint, lastNode *ListNode
	var index int
	curNode = head
	for curNode != nil {
		if index == pos {
			nodeToPoint = curNode
		}
		if curNode.Next == nil {
			lastNode = curNode
		}
		index++
		curNode = curNode.Next
	}
	if nodeToPoint != nil && lastNode != nil {
		lastNode.Next = nodeToPoint
	}
	return head
}
