package sorting

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		checkerFn checker
		want      int
	}{
		{
			name: "test1",
			args: args{n: 5},
			checkerFn: func(version int) bool {
				return version > 4
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultChecker = tt.checkerFn
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
