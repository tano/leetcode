package linked_list

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1,2,2,1",
			args: args{
				head: parseToListNodes("1,2,2,1"),
			},
			want: true,
		},
		{
			name: "1,2",
			args: args{
				head: parseToListNodes("1,2"),
			},
			want: false,
		},
		{
			name: "1",
			args: args{
				head: parseToListNodes("1"),
			},
			want: true,
		},
		{
			name: "0,0",
			args: args{
				head: parseToListNodes("0,0"),
			},
			want: true,
		},
		{
			name: "1,0,0",
			args: args{
				head: parseToListNodes("1,0,0"),
			},
			want: false,
		},
		{
			name: "1,0,1",
			args: args{
				head: parseToListNodes("1,0,1"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
