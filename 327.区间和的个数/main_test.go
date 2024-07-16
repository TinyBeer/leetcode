package main

import "testing"

func Test_countRangeSum(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [-2,5,-1], lower = -2, upper = 2",
			args: args{
				nums:  []int{-2, 5, -1},
				lower: -2,
				upper: 2,
			},
			want: 3,
		},
		{
			name: "nums = [0], lower = 0, upper = 0",
			args: args{
				nums:  []int{0},
				lower: 0,
				upper: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRangeSum(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countRangeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
