package array

import "testing"

func Test_convertToNumber(t *testing.T) {
	type args struct {
		sym byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				sym: '5',
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				sym: '.',
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToNumber(tt.args.sym)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertToNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				board: [][]rune{
					{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
					{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
					{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
					{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
					{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
					{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
					{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				board: [][]rune{
					{'.', '2', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '5', '.', '1'},
					{'.', '.', '.', '.', '.', '.', '8', '1', '3'},
					{'4', '.', '9', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '2', '.', '.', '.', '.', '.', '.'},
					{'7', '.', '6', '.', '.', '.', '.', '.', '.'},
					{'9', '.', '.', '.', '.', '4', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(stringsToBytes(tt.args.board)); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringsToBytes(strings [][]rune) [][]byte {
	result := make([][]byte, len(strings))
	for i := 0; i < len(strings); i++ {
		result[i] = make([]byte, len(strings))
		for j := 0; j < len(strings[i]); j++ {
			result[i][j] = byte(strings[i][j])
		}
	}
	return result
}
