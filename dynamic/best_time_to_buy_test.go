package dynamic

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{
					7, 1, 5, 3, 6, 4,
				},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				prices: []int{
					7, 6, 4, 3, 1,
				},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				prices: []int{
					1, 2,
				},
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				prices: []int{
					2, 4, 1,
				},
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				prices: []int{
					2, 1, 2, 0, 1,
				},
			},
			want: 1,
		},
		{
			name: "6",
			args: args{
				prices: []int{
					3, 3, 5, 0, 0, 3, 1, 4,
				},
			},
			want: 4,
		},
		{
			name: "7",
			args: args{
				prices: []int{
					3, 2, 6, 5, 0, 3,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
