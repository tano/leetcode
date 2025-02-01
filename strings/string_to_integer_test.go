package strings

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "2",
			args: args{
				s: " -042",
			},
			want: -42,
		},
		{
			name: "3",
			args: args{
				s: "1337c0d3",
			},
			want: 1337,
		},
		{
			name: "4",
			args: args{
				s: "0-1",
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				s: "words and 987",
			},
			want: 0,
		},
		{
			name: "6",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "7",
			args: args{
				s: "   -042",
			},
			want: -42,
		},
		{
			name: "8",
			args: args{
				s: "+1",
			},
			want: 1,
		},
		{
			name: "9",
			args: args{
				s: "+-12",
			},
			want: 0,
		},
		{
			name: "10",
			args: args{
				s: "20000000000000000000",
			},
			want: 2147483647,
		},
		{
			name: "11",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "12",
			args: args{
				s: "+",
			},
			want: 0,
		},
		{
			name: "13",
			args: args{
				s: "-13+8",
			},
			want: -13,
		},
		{
			name: "14",
			args: args{
				s: " ++1",
			},
			want: 0,
		},
		{
			name: "15",
			args: args{
				s: "-9223372036854775809",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
