package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,1,3,2,5]",
			args: args{
				nums: []int{1, 2, 1, 3, 2, 5},
			},
			want: []int{3, 5},
		},
		{
			name: "nums = [-1,0]",
			args: args{
				nums: []int{-1, 0},
			},
			want: []int{-1, 0},
		},
		{
			name: "nums = [0,1]",
			args: args{
				nums: []int{0, 1},
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.args.nums)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
