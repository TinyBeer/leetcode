package main

import "testing"

func Test_minPatches(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,3], n = 6",
			args: args{
				nums: []int{1, 3},
				n:    6,
			},
			want: 1,
		},
		{
			name: "nums = [1,5,10], n = 20",
			args: args{
				nums: []int{1, 5, 10},
				n:    20,
			},
			want: 2,
		},
		{
			name: "nums = [1,2,2], n = 5",
			args: args{
				nums: []int{1, 2, 2},
				n:    5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPatches(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("minPatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
