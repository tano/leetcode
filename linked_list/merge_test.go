package linked_list

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				list1: parseToListNodes("1,2,4"),
				list2: parseToListNodes("1,3,4"),
			},
			want: parseToListNodes("1,1,2,3,4,4"),
		},
		{
			name: "Example 2",
			args: args{
				list1: parseToListNodes(""),
				list2: parseToListNodes(""),
			},
			want: parseToListNodes(""),
		},
		{
			name: "Example 3",
			args: args{
				list1: parseToListNodes(""),
				list2: parseToListNodes("0"),
			},
			want: parseToListNodes("0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseToListNodes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1,2,4",
			args: args{
				s: "1,2,4",
			},
			want: func() *ListNode {
				first := ListNode{Val: 1}
				second := ListNode{Val: 2}
				third := ListNode{Val: 4}
				first.Next = &second
				second.Next = &third
				return &first
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseToListNodes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseToListNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
