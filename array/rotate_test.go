package array

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,3,4,5,6,7], k = 3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "nums = [-1,-100,3,99], k = 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		},
		{
			name: "nums = [-2], k = 2",
			args: args{
				nums: []int{-2},
				k:    2,
			},
			want: []int{-2},
		},
		{
			name: "nums = [1,2], k = 3",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			want: []int{2, 1},
		},
		{
			name: "nums = [1,2,3], k = 4",
			args: args{
				nums: []int{1, 2, 3},
				k:    4,
			},
			want: []int{3, 1, 2},
		},
		{
			name: "nums = [1,2,3,4,5,6], k = 11",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    11,
			},
			want: []int{2, 3, 4, 5, 6, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
