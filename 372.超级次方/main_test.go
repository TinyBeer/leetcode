package main

import "testing"

func Test_superPow(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a = 2, b = [3]",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "a = 2, b = [1,0]",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 10,
		},
		{
			name: "a = 1, b = [4,3,3,8,5,2]",
			args: args{
				a: 1,
				b: []int{4, 3, 3, 8, 5, 2},
			},
			want: 1,
		},
		{
			name: "a = 2147483647, b = [2,0,0]",
			args: args{
				a: 2147483647,
				b: []int{2, 0, 0},
			},
			want: 1198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
