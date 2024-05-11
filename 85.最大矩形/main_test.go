package main

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]`,
			args: args{
				matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
			},
			want: 6,
		},
		{
			name: `matrix = [["0"]]`,
			args: args{
				matrix: [][]byte{{'0'}},
			},
			want: 0,
		},
		{
			name: `matrix = [["1"]]`,
			args: args{
				matrix: [][]byte{{'1'}},
			},
			want: 1,
		},
		{
			name: `matrix = [["0","1"]]`,
			args: args{
				matrix: [][]byte{{'0', '1'}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
