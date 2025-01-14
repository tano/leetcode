package linked_list

import (
	"reflect"
	"testing"
)

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
		{
			name: "1",
			args: args{
				s: "1",
			},
			want: func() *ListNode {
				first := ListNode{Val: 1}
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
