package trees

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: func() *TreeNode {
				root := &TreeNode{Val: 0}
				root.Left = &TreeNode{Val: -3}
				root.Left.Left = &TreeNode{Val: -10}
				root.Right = &TreeNode{Val: 9}
				root.Right.Left = &TreeNode{Val: 5}
				return root
			}(),
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 3},
			},
			want: func() *TreeNode {
				root := &TreeNode{Val: 3}
				root.Left = &TreeNode{Val: 1}
				return root
			}(),
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 5, 8},
			},
			want: func() *TreeNode {
				root := &TreeNode{Val: 5}
				root.Left = &TreeNode{Val: 3}
				root.Right = &TreeNode{Val: 8}
				return root
			}(),
		},
		{
			name: "4",
			args: args{
				nums: []int{-1, 0, 1, 2},
			},
			want: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Left = &TreeNode{Val: 0}
				root.Left.Left = &TreeNode{Val: -1}
				root.Right = &TreeNode{Val: 2}
				return root
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
