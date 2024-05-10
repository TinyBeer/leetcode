package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nums = [2,5,6,0,0,1,2], target = 0",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			name: "nums = [2,5,6,0,0,1,2], target = 3",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 3,
			},
			want: false,
		},
		{
			name: "nums = [1,0,1,1,1], target = 0",
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
