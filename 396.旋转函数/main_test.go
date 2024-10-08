package main

import "testing"

func Test_maxRotateFunction(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [4,3,2,6]",
			args: args{
				nums: []int{4, 3, 2, 6},
			},
			want: 26,
		},
		{
			name: "nums = [100]",
			args: args{
				nums: []int{100},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRotateFunction(tt.args.nums); got != tt.want {
				t.Errorf("maxRotateFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
