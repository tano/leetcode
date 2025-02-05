package linked_list

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: parseToListNodes("1,2,3,4,5"),
				n:    2,
			},
			want: parseToListNodes("1,2,3,5"),
		},
		{
			name: "test2",
			args: args{
				head: parseToListNodes("1"),
				n:    1,
			},
			want: parseToListNodes(""),
		},
		{
			name: "test3",
			args: args{
				head: parseToListNodes("1,2"),
				n:    1,
			},
			want: parseToListNodes("1"),
		},
		{
			name: "test4",
			args: args{
				head: parseToListNodes("1,2"),
				n:    2,
			},
			want: parseToListNodes("2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %s, want %s", got, tt.want)
			}
		})
	}
}
