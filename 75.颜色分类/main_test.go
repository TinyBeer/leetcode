package main

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [2,0,2,1,1,0]",
			args: args{nums: []int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "nums = [2,0,1]",
			args: args{nums: []int{2, 0, 1}},
			want: []int{0, 1, 2},
		},
		{
			name: "nums = [1,2,0,0]",
			args: args{nums: []int{1, 2, 0, 0}},
			want: []int{0, 0, 1, 2},
		},
		{
			name: "nums = [0,1,2,0,0,2,2,1]",
			args: args{nums: []int{0, 1, 2, 0, 0, 2, 2, 1}},
			want: []int{0, 0, 0, 1, 1, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortColors()  %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

// 0,1,2,0,0,2,2,1
//
