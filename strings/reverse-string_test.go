package strings

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "test1",
			args: args{
				s: []rune{'h', 'e', 'l', 'l', 'o'},
			},
			want: []rune{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "test2",
			args: args{
				s: []rune{'H', 'a', 'n', 'n', 'a', 'h'},
			},
			want: []rune{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := runesToBytes(tt.args.s)
			reverseString(data)
			want := runesToBytes(tt.want)
			if !reflect.DeepEqual(data, want) {
				t.Errorf("reverseString() = %v, want %v", data, tt.want)
			}
		})
	}
}

func runesToBytes(runes []rune) []byte {
	result := make([]byte, len(runes))
	for i := 0; i < len(runes); i++ {
		result[i] = byte(runes[i])
	}
	return result
}
