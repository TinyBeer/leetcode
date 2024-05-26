package main

import "testing"

func Test_maximumGap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [3,6,9,1]",
			args: args{
				nums: []int{3, 6, 9, 1},
			},
			want: 3,
		},
		{
			name: "nums = [10]",
			args: args{
				nums: []int{10},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGap(tt.args.nums); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
