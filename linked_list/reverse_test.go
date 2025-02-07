package linked_list

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
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
			},
			want: parseToListNodes("5,4,3,2,1"),
		},
		{
			name: "test2",
			args: args{
				head: parseToListNodes("1,2"),
			},
			want: parseToListNodes("2,1"),
		},
		{
			name: "test3",
			args: args{
				head: parseToListNodes(""),
			},
			want: parseToListNodes(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
