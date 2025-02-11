package trees

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				root: func() *TreeNode {
					rootNode := &TreeNode{Val: 3}
					rootNode.Left = &TreeNode{Val: 9}
					rootNode.Right = &TreeNode{Val: 20}
					rootNode.Right.Left = &TreeNode{Val: 15}
					rootNode.Right.Right = &TreeNode{Val: 7}
					return rootNode
				}(),
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "2",
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
