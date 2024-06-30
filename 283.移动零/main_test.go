package main

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [0,1,0,3,12]",
			args: args{nums: []int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "nums = [0]",
			args: args{nums: []int{0}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() got: %v, expect: %v\n", tt.args.nums, tt.want)
			}
		})
	}
}
