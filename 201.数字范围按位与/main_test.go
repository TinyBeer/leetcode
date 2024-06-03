package main

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "left = 5, right = 7",
			args: args{
				left:  5,
				right: 7,
			},
			want: 4,
		},
		{
			name: "left = 0, right = 0",
			args: args{
				left:  0,
				right: 0,
			},
			want: 0,
		},
		{
			name: "left = 1, right = 2147483647",
			args: args{
				left:  1,
				right: 2147483647,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
