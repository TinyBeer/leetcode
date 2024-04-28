package main

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [-1,2,1,-4], target = 1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "nums = [0,0,0], target = 1",
			args: args{
				nums:   []int{0, 0, 0},
				target: 0,
			},
			want: 0,
		},
		{
			name: "nums = [1,1,1,1], target = 0",
			args: args{
				nums:   []int{1, 1, 1, 1},
				target: 0,
			},
			want: 3,
		},
		{
			name: "nums = [0,1,2], target = 3",
			args: args{
				nums:   []int{0, 1, 2},
				target: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
