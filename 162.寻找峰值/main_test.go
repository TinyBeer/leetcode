package main

import (
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,3,1]",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: []int{2},
		},
		{
			name: "nums = [1,2]",
			args: args{
				nums: []int{1, 2},
			},
			want: []int{1},
		},
		{
			name: "nums = [1,2,1,3,5,6,4]",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: []int{1, 5},
		},
		{
			name: "nums = [3,1,2]",
			args: args{
				nums: []int{3, 1, 2},
			},
			want: []int{0, 2},
		},
		{
			name: "nums = [1,2,3]",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.args.nums)
			flag := false
			for _, v := range tt.want {
				if v == got {
					flag = true
					break
				}
			}
			if flag == false {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
